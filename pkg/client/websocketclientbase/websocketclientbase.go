package websocketclientbase

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/huobirdcenter/huobi_golang/internal/gzip"
	"github.com/huobirdcenter/huobi_golang/internal/model"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"strings"
	"sync"
	"time"
)

const (
	TimerIntervalSecond = 5
	ReconnectWaitSecond = 60

	wsPath = "/ws"
	feedPath = "/feed"
)

// It will be invoked after websocket connected
type ConnectedHandler func()

// It will be invoked after valid message received
type MessageHandler func(message string) (interface{}, error)

// It will be invoked after response is parsed
type ResponseHandler func(response interface{})

// The base class that responsible to get data from websocket
type WebSocketClientBase struct {
	host              string
	path              string
	conn              *websocket.Conn
	connectedHandler  ConnectedHandler
	messageHandler    MessageHandler
	responseHandler   ResponseHandler
	stopReadChannel   chan int
	stopTickerChannel chan int
	ticker            *time.Ticker
	lastReceivedTime  time.Time
	sendMutex         *sync.Mutex
}

// Initializer
func (p *WebSocketClientBase) Init(host string) *WebSocketClientBase {
	p.host = host
	p.path = wsPath
	p.stopReadChannel = make(chan int, 1)
	p.stopTickerChannel = make(chan int, 1)
	p.sendMutex = &sync.Mutex{}

	return p
}

// Initializer with path
func (p *WebSocketClientBase) InitWithFeedPath(host string) *WebSocketClientBase {
	p.Init(host)
	p.path = feedPath
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
func (p *WebSocketClientBase) Connect(autoConnect bool) {
	// initialize last received time as now
	p.lastReceivedTime = time.Now()

	// connect to websocket
	p.connectWebSocket()

	// start ticker to manage connection
	if autoConnect {
		p.startTicker()
	}
}

// Send data to websocket server
func (p *WebSocketClientBase) Send(data string) {
	if p.conn == nil {
		applogger.Error("WebSocket sent error: no connection available")
		return
	}

	p.sendMutex.Lock()
	err := p.conn.WriteMessage(websocket.TextMessage, []byte(data))
	p.sendMutex.Unlock()

	if err != nil {
		applogger.Error("WebSocket sent error: data=%s, error=%s", data, err)
	}
}

// Close the connection to server
func (p *WebSocketClientBase) Close() {
	p.stopTicker()
	p.disconnectWebSocket()
}

// connect to server
func (p *WebSocketClientBase) connectWebSocket() {
	var err error
	url := fmt.Sprintf("wss://%s%s", p.host, p.path)
	applogger.Debug("WebSocket connecting...")
	p.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		applogger.Error("WebSocket connected error: %s", err)
		return
	}
	applogger.Info("WebSocket connected")

	// start loop to read and handle message
	p.startReadLoop()

	if p.connectedHandler != nil {
		p.connectedHandler()
	}
}

// disconnect with server
func (p *WebSocketClientBase) disconnectWebSocket() {
	if p.conn == nil {
		return
	}

	// start a new goroutine to send stop signal
	go p.stopReadLoop()

	applogger.Debug("WebSocket disconnecting...")
	err := p.conn.Close()
	if err != nil {
		applogger.Error("WebSocket disconnect error: %s", err)
		return
	}

	applogger.Info("WebSocket disconnected")
}

// initialize a ticker and start a goroutine tickerLoop()
func (p *WebSocketClientBase) startTicker() {
	p.ticker = time.NewTicker(TimerIntervalSecond * time.Second)

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
	applogger.Debug("tickerLoop started")
	for {
		select {
		// Receive data from stopChannel
		case <-p.stopTickerChannel:
			applogger.Debug("tickerLoop stopped")
			return

		// Receive tick from tickChannel
		case <-p.ticker.C:
			elapsedSecond := time.Now().Sub(p.lastReceivedTime).Seconds()
			applogger.Debug("WebSocket received data %f sec ago", elapsedSecond)

			if elapsedSecond > ReconnectWaitSecond {
				applogger.Info("WebSocket reconnect...")
				p.disconnectWebSocket()
				p.connectWebSocket()
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
	applogger.Debug("readLoop started")
	for {
		select {
		// Receive data from stopChannel
		case <-p.stopReadChannel:
			applogger.Debug("readLoop stopped")
			return

		default:
			if p.conn == nil {
				applogger.Error("Read error: no connection available")
				time.Sleep(TimerIntervalSecond * time.Second)
				continue
			}

			msgType, buf, err := p.conn.ReadMessage()
			if err != nil {
				applogger.Error("Read error: %s", err)
				time.Sleep(TimerIntervalSecond * time.Second)
				continue
			}

			p.lastReceivedTime = time.Now()

			// decompress gzip data if it is binary message
			if msgType == websocket.BinaryMessage {
				message, err := gzip.GZipDecompress(buf)
				if err != nil {
					applogger.Error("UnGZip data error: %s", err)
					return
				}

				// Try to pass as PingMessage
				pingMsg := model.ParsePingMessage(message)

				// If it is Ping then respond Pong
				if pingMsg != nil && pingMsg.Ping != 0 {
					applogger.Debug("Received Ping: %d", pingMsg.Ping)
					pongMsg := fmt.Sprintf("{\"pong\": %d}", pingMsg.Ping)
					p.Send(pongMsg)
					applogger.Debug("Replied Pong: %d", pingMsg.Ping)
				} else if strings.Contains(message, "tick") || strings.Contains(message, "data") {
					// If it contains expected string, then invoke message handler and response handler
					result, err := p.messageHandler(message)
					if err != nil {
						applogger.Error("Handle message error: %s", err)
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
