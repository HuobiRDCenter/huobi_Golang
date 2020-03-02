package marketwebsocketclient

import (
	"../../response/market"
	"../websocketclientbase"
	"encoding/json"
	"fmt"
)

type DepthWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

func (p *DepthWebSocketClient) Init(host string) *DepthWebSocketClient {
	p.WebSocketClientBase.Init(host)
	return p
}

func (p *DepthWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

func (p *DepthWebSocketClient) Request(symbol string, step string, clientId string) error {
	req := fmt.Sprintf("{\"req\": \"market.%s.depth.%s\",\"id\": \"%s\" }", symbol, step, clientId)
	return p.WebSocketClientBase.Send(req)
}

func (p *DepthWebSocketClient) Subscribe(symbol string, step string, clientId string) error {
	sub := fmt.Sprintf("{\"sub\": \"market.%s.depth.%s\",\"id\": \"%s\" }", symbol, step, clientId)
	return p.WebSocketClientBase.Send(sub)
}

func (p *DepthWebSocketClient) UnSubscribe(symbol string, step string, clientId string) error {
	unsub := fmt.Sprintf("{\"unsub\": \"market.%s.depth.%s\",\"id\": \"%s\" }", symbol, step, clientId)
	return p.WebSocketClientBase.Send(unsub)
}

func (p *DepthWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeDepthResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
