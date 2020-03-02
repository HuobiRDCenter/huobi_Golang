package marketwebsocketclient

import (
	"../../response/market"
	"../websocketclientbase"
	"encoding/json"
	"fmt"
)

type Last24hCandlestickWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

func (p *Last24hCandlestickWebSocketClient) Init(host string) *Last24hCandlestickWebSocketClient {
	p.WebSocketClientBase.Init(host)
	return p
}

func (p *Last24hCandlestickWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

func (p *Last24hCandlestickWebSocketClient) Request(symbol string, clientId string) error {
	req := fmt.Sprintf("{\"req\": \"market.%s.detail\",\"id\": \"%s\" }", symbol, clientId)
	return p.WebSocketClientBase.Send(req)
}

func (p *Last24hCandlestickWebSocketClient) Subscribe(symbol string, clientId string) error {
	sub := fmt.Sprintf("{\"sub\": \"market.%s.detail\",\"id\": \"%s\" }", symbol, clientId)
	return p.WebSocketClientBase.Send(sub)
}

func (p *Last24hCandlestickWebSocketClient) UnSubscribe(symbol string, clientId string) error {
	unsub := fmt.Sprintf("{\"unsub\": \"market.%s.detail\",\"id\": \"%s\" }", symbol, clientId)
	return p.WebSocketClientBase.Send(unsub)
}

func (p *Last24hCandlestickWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeLast24hCandlestickResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
