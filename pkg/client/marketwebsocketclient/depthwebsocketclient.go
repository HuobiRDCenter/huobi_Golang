package marketwebsocketclient

import (
	"encoding/json"
	"fmt"
	"github.com/newgoo/huobi_golang/logging/applogger"
	"github.com/newgoo/huobi_golang/pkg/client/websocketclientbase"
	"github.com/newgoo/huobi_golang/pkg/model/market"
)

// Responsible to handle Depth data from WebSocket
type DepthWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

// Initializer
func (p *DepthWebSocketClient) Init(host string) *DepthWebSocketClient {
	p.WebSocketClientBase.Init(host)
	return p
}

// Set callback handler
func (p *DepthWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Request full depth data
func (p *DepthWebSocketClient) Request(symbol string, step string, clientId string) {
	topic := fmt.Sprintf("market.%s.depth.%s", symbol, step)
	req := fmt.Sprintf("{\"req\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.WebSocketClientBase.Send(req)

	applogger.Info("WebSocket requested, topic=%s, clientId=%s", topic, clientId)
}

// Subscribe latest market by price order book in snapshot mode at 1-second interval.
func (p *DepthWebSocketClient) Subscribe(symbol string, step string, clientId string) {
	topic := fmt.Sprintf("market.%s.depth.%s", symbol, step)
	sub := fmt.Sprintf("{\"sub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe market by price order book
func (p *DepthWebSocketClient) UnSubscribe(symbol string, step string, clientId string) {
	topic := fmt.Sprintf("market.%s.depth.%s", symbol, step)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *DepthWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeDepthResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
