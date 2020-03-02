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
	websocketV1Path = "/ws/v1"
)

type AuthenticationV1ResponseHandler func(resp *model.WebSocketV1AuthenticationResponse)

type WebSocketV1ClientBase struct {
	host string
	conn *websocket.Conn

	authenticationResponseHandler AuthenticationV1ResponseHandler
	messageHandler                MessageHandler
	responseHandler               ResponseHandler

	stopReadChannel   chan int
	stopTickerChannel chan int
	ticker            *time.Ticker
	lastReceivedTime  time.Time

	requestBuilder *requestbuilder.WebSocketV1RequestBuilder
}

func (p *WebSocketV1ClientBase) Init(accessKey string, secretKey string, host string) *WebSocketV1ClientBase {
	p.host = host
	p.stopReadChannel = make(chan int, 1)
	p.stopTickerChannel = make(chan int, 1)
	p.requestBuilder = new(requestbuilder.WebSocketV1RequestBuilder).Init(accessKey, secretKey, host, websocketV1Path)
	return p
}

func (p *WebSocketV1ClientBase) SetHandler(authHandler AuthenticationV1ResponseHandler, msgHandler MessageHandler, repHandler ResponseHandler) {
	p.authenticationResponseHandler = authHandler
	p.messageHandler = msgHandler
	p.responseHandler = repHandler
}

func (p *WebSocketV1ClientBase) Connect(autoConnect bool) error {
	err := p.connectWebSocket()
	if err != nil {
		return err
	}

	if autoConnect {
		p.startTicker()
	}

	return nil
}

func (p *WebSocketV1ClientBase) Send(data string) error {
	if p.conn == nil {
		return errors.New("no connection available")
	}

	err := p.conn.WriteMessage(websocket.TextMessage, []byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (p *WebSocketV1ClientBase) Close() {
	p.stopTicker()
	p.disconnectWebSocket()
}

func (p *WebSocketV1ClientBase) connectWebSocket() error {
	var err error
	url := fmt.Sprintf("wss://%s%s", p.host, websocketV1Path)
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

func (p *WebSocketV1ClientBase) disconnectWebSocket() {
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

func (p *WebSocketV1ClientBase) startTicker() {
	p.ticker = time.NewTicker(TimerIntervalSecond * time.Second)
	p.lastReceivedTime = time.Now()

	go p.tickerLoop()
}

func (p *WebSocketV1ClientBase) stopTicker() {
	if p.ticker != nil {
		p.ticker.Stop()
	}
	p.stopTickerChannel <- 1
}

func (p *WebSocketV1ClientBase) tickerLoop() {
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

func (p *WebSocketV1ClientBase) startReadLoop() {
	go p.readLoop()
}

func (p *WebSocketV1ClientBase) stopReadLoop() {
	p.stopReadChannel <- 1
}

func (p *WebSocketV1ClientBase) readLoop() {
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

			if msgType == websocket.BinaryMessage {
				message, err := gzip.GZipDecompress(buf)
				if err != nil {
					fmt.Printf("UnGZip data error: %s\n", err)
				}

				pingV1Msg := model.ParsePingV1Message(message)
				if pingV1Msg.IsPing() {
					fmt.Printf("Received Ping: %d\n", pingV1Msg.Timestamp)
					pongMsg := fmt.Sprintf("{\"op\": \"pong\", \"ts\": %d}", pingV1Msg.Timestamp)
					p.Send(pongMsg)
					fmt.Printf("Respond  Pong: %d\n", pingV1Msg.Timestamp)
				} else {
					authResp := model.ParseWSV1AuthResp(message)
					if authResp != nil && authResp.IsAuth() {
						if p.authenticationResponseHandler != nil {
							p.authenticationResponseHandler(authResp)
						}
					} else if strings.Contains(message, "balance") { //TODO: should use better way to determine
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
}
