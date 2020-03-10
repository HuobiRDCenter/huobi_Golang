package marketwebsocketclient

import (
	"encoding/json"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/pkg/client/websocketclientbase"
	"github.com/huobirdcenter/huobi_golang/pkg/response/market"
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
func (p *DepthWebSocketClient) Request(symbol string, step string, clientId string) error {
	req := fmt.Sprintf("{\"req\": \"market.%s.depth.%s\",\"id\": \"%s\" }", symbol, step, clientId)
	return p.WebSocketClientBase.Send(req)
}

// Subscribe latest market by price order book in snapshot mode at 1-second interval.
func (p *DepthWebSocketClient) Subscribe(symbol string, step string, clientId string) error {
	sub := fmt.Sprintf("{\"sub\": \"market.%s.depth.%s\",\"id\": \"%s\" }", symbol, step, clientId)
	return p.WebSocketClientBase.Send(sub)
}

// Unsubscribe market by price order book
func (p *DepthWebSocketClient) UnSubscribe(symbol string, step string, clientId string) error {
	unsub := fmt.Sprintf("{\"unsub\": \"market.%s.depth.%s\",\"id\": \"%s\" }", symbol, step, clientId)
	return p.WebSocketClientBase.Send(unsub)
}

func (p *DepthWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeDepthResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
