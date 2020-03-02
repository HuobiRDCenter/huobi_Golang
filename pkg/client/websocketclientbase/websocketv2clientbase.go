package websocketclientbase

import (
	"../../../internal/gzip"
	"../../../internal/model"
	"../../../internal/requestbuilder"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"strings"
	"time"
)

const (
	websocketV2Path = "/ws/v2"
)

type AuthenticationV2ResponseHandler func(resp *model.WebSocketV2AuthenticationResponse)

type WebSocketV2ClientBase struct {
	host                     string
	conn                     *websocket.Conn

	authenticationResponseHandler AuthenticationV2ResponseHandler
	messageHandler                MessageHandler
	responseHandler               ResponseHandler

	stopReadChannel          chan int
	stopTickerChannel        chan int
	ticker                   *time.Ticker
	lastReceivedTime         time.Time

	requestBuilder *requestbuilder.WebSocketV2RequestBuilder
}

func (p *WebSocketV2ClientBase) Init(accessKey string, secretKey string, host string) *WebSocketV2ClientBase {
	p.host = host
	p.stopReadChannel = make(chan int, 1)
	p.stopTickerChannel = make(chan int, 1)
	p.requestBuilder = new(requestbuilder.WebSocketV2RequestBuilder).Init(accessKey, secretKey, host, websocketV2Path)
	return p
}

func (p *WebSocketV2ClientBase) SetHandler(authHandler AuthenticationV2ResponseHandler, msgHandler MessageHandler, repHandler ResponseHandler) {
	p.authenticationResponseHandler = authHandler
	p.messageHandler = msgHandler
	p.responseHandler = repHandler
}

func (p *WebSocketV2ClientBase) Connect(autoConnect bool) error {
	err := p.connectWebSocket()
	if err != nil {
		return err
	}

	if autoConnect {
		p.startTicker()
	}

	return nil
}

func (p *WebSocketV2ClientBase) Send(data string) error {
	if p.conn == nil {
		return errors.New("no connection available")
	}

	err := p.conn.WriteMessage(websocket.TextMessage, []byte(data))
	if err != nil {
		return err
	}

	return nil
}

func (p *WebSocketV2ClientBase) Close() {
	p.stopTicker()
	p.disconnectWebSocket()
}

func (p *WebSocketV2ClientBase) connectWebSocket() error {
	var err error
	url := fmt.Sprintf("wss://%s%s", p.host, websocketV2Path)
	fmt.Println("WebSocket connecting...")
	p.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}
	fmt.Println("WebSocket connected")

	auth, err := p.requestBuilder.Build()
	if err != nil {
		return err
	}
	err = p.Send(auth)
	if err != nil {
		return err
	}

	p.startReadLoop()

	return nil
}

func (p *WebSocketV2ClientBase) disconnectWebSocket() {
	if p.conn == nil {
		return
	}

	p.stopReadLoop()

	fmt.Println("WebSocket disconnecting...")
	err := p.conn.Close()
	if err != nil {
		fmt.Printf("WebSocket disconnect error: %s\n", err)
		return
	}

	fmt.Println("WebSocket disconnected")
}

func (p *WebSocketV2ClientBase) startTicker() {
	p.ticker = time.NewTicker(TimerIntervalSecond * time.Second)
	p.lastReceivedTime = time.Now()

	go p.tickerLoop()
}

func (p *WebSocketV2ClientBase) stopTicker() {
	p.ticker.Stop()
	p.stopTickerChannel <- 1
}

func (p *WebSocketV2ClientBase) tickerLoop() {
	fmt.Println("tickerLoop started")
	for {
		select {
		case <-p.stopTickerChannel:
			fmt.Println("tickerLoop stopped")
			return

		case <-p.ticker.C:
			elapsedSecond := time.Now().Sub(p.lastReceivedTime).Seconds()
			fmt.Printf("WebSocket received data %f sec ago\n", elapsedSecond)

			if elapsedSecond > ReconnectWaitSecond {
				fmt.Println("WebSocket reconnect...")
				p.disconnectWebSocket()
				err := p.connectWebSocket()
				if err != nil {
					fmt.Printf("WebSocket reconnect error: %s\n", err)
				}
			}
		}
	}
}

func (p *WebSocketV2ClientBase) startReadLoop() {
	go p.readLoop()
}

func (p *WebSocketV2ClientBase) stopReadLoop() {
	p.stopReadChannel <- 1 //TODO: consider put this into goroutine to unblock
}

func (p *WebSocketV2ClientBase) readLoop() {
	fmt.Println("readLoop started")
	for {
		select {
		case <-p.stopReadChannel:
			fmt.Println("readLoop stopped")
			return

		default:
			if p.conn == nil {
				fmt.Printf("Read error: no connection available")
				time.Sleep(TimerIntervalSecond * time.Second)
				continue
			}

			msgType, buf, err := p.conn.ReadMessage()
			if err != nil {
				fmt.Printf("Read error: %s\n", err)
				time.Sleep(TimerIntervalSecond * time.Second)
				continue
			}

			p.lastReceivedTime = time.Now()

			var message string
			if msgType == websocket.BinaryMessage {
				message, err = gzip.GZipDecompress(buf)
				if err != nil {
					fmt.Printf("UnGZip data error: %s\n", err)
				}
			} else if msgType == websocket.TextMessage {
				message = string(buf)
			}

			pingV2Msg := model.ParsePingV2Message(message)
			if pingV2Msg.IsPing() {
				fmt.Printf("Received Ping: %d\n", pingV2Msg.Data.Timestamp)
				pongMsg := fmt.Sprintf("{\"action\": \"pong\", \"data\": { \"ts\": %d } }", pingV2Msg.Data.Timestamp)
				p.Send(pongMsg)
				fmt.Printf("Respond  Pong: %d\n", pingV2Msg.Data.Timestamp)
			} else {
				authResp := model.ParseWSV2AuthResp(message)
				if authResp != nil && authResp.IsAuth() {
					if p.authenticationResponseHandler != nil {
						p.authenticationResponseHandler(authResp)
					}
				} else if strings.Contains(message, "symbol") || strings.Contains(message, "currency") {//TODO: should use better way to determine
					result, err := p.messageHandler(message)
					if err != nil {
						fmt.Printf("Handle message error: %s\n", err)
						continue
					}
					if p.responseHandler != nil {
						p.responseHandler(result)
					}
				}
			}
		}
	}
}
