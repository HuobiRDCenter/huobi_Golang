package websocketclientbase

import (
	"../../../internal/gzip"
	"../../../internal/model"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"strings"
	"time"
)

const (
	TimerIntervalSecond = 5
	ReconnectWaitSecond = 60

	path = "/ws"
)

type ConnectedHandler func()

type MessageHandler func(message string) (interface{}, error)

type ResponseHandler func(resp interface{})

type WebSocketClientBase struct {
	host              string
	conn              *websocket.Conn
	connectedHandler  ConnectedHandler
	messageHandler    MessageHandler
	responseHandler   ResponseHandler
	stopReadChannel   chan int
	stopTickerChannel chan int
	ticker            *time.Ticker
	lastReceivedTime  time.Time
}

func (p *WebSocketClientBase) Init(host string) *WebSocketClientBase {
	p.host = host
	p.stopReadChannel = make(chan int, 1)
	p.stopTickerChannel = make(chan int, 1)
	return p
}

func (p *WebSocketClientBase) SetHandler(connHandler ConnectedHandler, msgHandler MessageHandler, repHandler ResponseHandler) {
	p.connectedHandler = connHandler
	p.messageHandler = msgHandler
	p.responseHandler = repHandler
}

func (p *WebSocketClientBase) Connect(autoConnect bool) error {
	err := p.connectWebSocket()
	if err != nil {
		return err
	}

	if autoConnect {
		p.startTicker()
	}

	return nil
}

func (p *WebSocketClientBase) Send(data string) error {
	if p.conn == nil {
		return errors.New("no connection available")
	}

	err := p.conn.WriteMessage(websocket.TextMessage, []byte(data))
	if err != nil {
		return err
	}

	return nil
}

func (p *WebSocketClientBase) Close() {
	p.stopTicker()
	p.disconnectWebSocket()
}

func (p *WebSocketClientBase) connectWebSocket() error {
	var err error
	url := fmt.Sprintf("wss://%s%s", p.host, path)
	fmt.Println("WebSocket connecting...")
	p.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}
	fmt.Println("WebSocket connected")

	p.startReadLoop()

	if p.connectedHandler != nil {
		p.connectedHandler()
	}

	return nil
}

func (p *WebSocketClientBase) disconnectWebSocket() {
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

func (p *WebSocketClientBase) startTicker() {
	p.ticker = time.NewTicker(TimerIntervalSecond * time.Second)
	p.lastReceivedTime = time.Now()

	go p.tickerLoop()
}

func (p *WebSocketClientBase) stopTicker() {
	p.ticker.Stop()
	p.stopTickerChannel <- 1
}

func (p *WebSocketClientBase) tickerLoop() {
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

func (p *WebSocketClientBase) startReadLoop() {
	go p.readLoop()
}

func (p *WebSocketClientBase) stopReadLoop() {
	p.stopReadChannel <- 1
}

func (p *WebSocketClientBase) readLoop() {
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

				pingMsg := model.ParsePingMessage(message)
				if pingMsg != nil && pingMsg.Ping != 0 {
					fmt.Printf("Received Ping: %d\n", pingMsg.Ping)
					pongMsg := fmt.Sprintf("{\"pong\": %d}", pingMsg.Ping)
					p.conn.WriteMessage(websocket.TextMessage, []byte(pongMsg))
					fmt.Printf("Replied Pong: %d\n", pingMsg.Ping)
				} else if strings.Contains(message, "tick") || strings.Contains(message, "data") { //TODO: should use better way to determine
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
