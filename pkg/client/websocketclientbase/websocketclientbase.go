package websocketclientbase

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/huobirdcenter/huobi_golang/internal/gzip"
	"github.com/huobirdcenter/huobi_golang/internal/model"
	"strings"
	"time"
)

const (
	TimerIntervalSecond = 5
	ReconnectWaitSecond = 60

	path = "/ws"
)

// It will be invoked after websocket connected
type ConnectedHandler func()

// It will be invoked after valid message received
type MessageHandler func(message string) (interface{}, error)

// It will be invoked after response is parsed
type ResponseHandler func(resp interface{})

// The base class that responsible to get data from websocket
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

// Initializer
func (p *WebSocketClientBase) Init(host string) *WebSocketClientBase {
	p.host = host
	p.stopReadChannel = make(chan int, 1)
	p.stopTickerChannel = make(chan int, 1)
	return p
}

// Set callback handler
func (p *WebSocketClientBase) SetHandler(connHandler ConnectedHandler, msgHandler MessageHandler, repHandler ResponseHandler) {
	p.connectedHandler = connHandler
	p.messageHandler = msgHandler
	p.responseHandler = repHandler
}

// Connect to websocket server
// if autoConnect is true, then the connection can be re-connect if no data received after the pre-defined timeout
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

// Send data to websocket server
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

// Close the connection to server
func (p *WebSocketClientBase) Close() {
	p.stopTicker()
	p.disconnectWebSocket()
}

// connect to server
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

// disconnect with server
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

// initialize a ticker and start a goroutine tickerLoop()
func (p *WebSocketClientBase) startTicker() {
	p.ticker = time.NewTicker(TimerIntervalSecond * time.Second)
	p.lastReceivedTime = time.Now()

	go p.tickerLoop()
}

// stop ticker and stop the goroutine
func (p *WebSocketClientBase) stopTicker() {
	p.ticker.Stop()
	p.stopTickerChannel <- 1
}

// defines a for loop that will run based on ticker's frequency
// It checks the last data that received from server, if it is longer than the threshold,
// it will force disconnect server and connect again.
func (p *WebSocketClientBase) tickerLoop() {
	for {
		select {
		// Receive data from stopChannel
		case <-p.stopTickerChannel:
			fmt.Println("tickerLoop stopped")
			return

		// Receive tick from tickChannel
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

// start a goroutine readLoop()
func (p *WebSocketClientBase) startReadLoop() {
	go p.readLoop()
}

// stop the goroutine readLoop()
func (p *WebSocketClientBase) stopReadLoop() {
	p.stopReadChannel <- 1
}

// defines a for loop to read data from server
// it will stop once it receives the signal from stopReadChannel
func (p *WebSocketClientBase) readLoop() {
	for {
		select {
		// Receive data from stopChannel
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

			// decompress gzip data if it is binary message
			if msgType == websocket.BinaryMessage {
				message, err := gzip.GZipDecompress(buf)
				if err != nil {
					fmt.Printf("UnGZip data error: %s\n", err)
				}

				// Try to pass as PingMessage
				pingMsg := model.ParsePingMessage(message)

				// If it is Ping then respond Pong
				if pingMsg != nil && pingMsg.Ping != 0 {
					fmt.Printf("Received Ping: %d\n", pingMsg.Ping)
					pongMsg := fmt.Sprintf("{\"pong\": %d}", pingMsg.Ping)
					p.conn.WriteMessage(websocket.TextMessage, []byte(pongMsg))
					fmt.Printf("Replied Pong: %d\n", pingMsg.Ping)
				} else if strings.Contains(message, "tick") || strings.Contains(message, "data") {
					// If it contains expected string, then invoke message handler and response handler
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
