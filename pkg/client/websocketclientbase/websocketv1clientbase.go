package websocketclientbase

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/huobirdcenter/huobi_golang/internal/gzip"
	"github.com/huobirdcenter/huobi_golang/internal/model"
	"github.com/huobirdcenter/huobi_golang/internal/requestbuilder"
	"strings"
	"time"
)

const (
	websocketV1Path = "/ws/v1"
)

// It will be invoked after websocket v1 authentication response received
type AuthenticationV1ResponseHandler func(resp *model.WebSocketV1AuthenticationResponse)

// The base class that responsible to get data from websocket authentication v1
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

// Initializer
func (p *WebSocketV1ClientBase) Init(accessKey string, secretKey string, host string) *WebSocketV1ClientBase {
	p.host = host
	p.stopReadChannel = make(chan int, 1)
	p.stopTickerChannel = make(chan int, 1)
	p.requestBuilder = new(requestbuilder.WebSocketV1RequestBuilder).Init(accessKey, secretKey, host, websocketV1Path)
	return p
}

// Set callback handler
func (p *WebSocketV1ClientBase) SetHandler(authHandler AuthenticationV1ResponseHandler, msgHandler MessageHandler, repHandler ResponseHandler) {
	p.authenticationResponseHandler = authHandler
	p.messageHandler = msgHandler
	p.responseHandler = repHandler
}

// Connect to websocket server
// if autoConnect is true, then the connection can be re-connect if no data received after the pre-defined timeout
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

// Send data to websocket server
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

// Close the connection to server
func (p *WebSocketV1ClientBase) Close() {
	p.stopTicker()
	p.disconnectWebSocket()
}

// connect to server
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

// disconnect with server
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

// initialize a ticker and start a goroutine tickerLoop()
func (p *WebSocketV1ClientBase) startTicker() {
	p.ticker = time.NewTicker(TimerIntervalSecond * time.Second)
	p.lastReceivedTime = time.Now()

	go p.tickerLoop()
}

// stop ticker and stop the goroutine
func (p *WebSocketV1ClientBase) stopTicker() {
	if p.ticker != nil {
		p.ticker.Stop()
	}
	p.stopTickerChannel <- 1
}

// defines a for loop that will run based on ticker's frequency
// It checks the last data that received from server, if it is longer than the threshold,
// it will force disconnect server and connect again.
func (p *WebSocketV1ClientBase) tickerLoop() {
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
func (p *WebSocketV1ClientBase) startReadLoop() {
	go p.readLoop()
}

// stop the goroutine readLoop()
func (p *WebSocketV1ClientBase) stopReadLoop() {
	p.stopReadChannel <- 1
}

// defines a for loop to read data from server
// it will stop once it receives the signal from stopReadChannel
func (p *WebSocketV1ClientBase) readLoop() {
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

				// Try to pass as PingV1Message
				// If it is Ping then respond Pong
				pingV1Msg := model.ParsePingV1Message(message)
				if pingV1Msg.IsPing() {
					fmt.Printf("Received Ping: %d\n", pingV1Msg.Timestamp)
					pongMsg := fmt.Sprintf("{\"op\": \"pong\", \"ts\": %d}", pingV1Msg.Timestamp)
					p.Send(pongMsg)
					fmt.Printf("Respond  Pong: %d\n", pingV1Msg.Timestamp)
				} else {
					// Try to pass as websocket v1 authentication response
					// If it is then invoke authentication handler
					authResp := model.ParseWSV1AuthResp(message)
					if authResp != nil && authResp.IsAuth() {
						if p.authenticationResponseHandler != nil {
							p.authenticationResponseHandler(authResp)
						}
					} else if strings.Contains(message, "balance") {
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
}
