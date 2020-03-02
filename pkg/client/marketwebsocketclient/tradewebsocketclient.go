package marketwebsocketclient

import (
	"../../response/market"
	"../websocketclientbase"
	"encoding/json"
	"fmt"
)

type TradeWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

func (p *TradeWebSocketClient) Init(host string) *TradeWebSocketClient {
	p.WebSocketClientBase.Init(host)
	return p
}

func (p *TradeWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

func (p *TradeWebSocketClient) Request(symbol string, clientId string) error {
	req := fmt.Sprintf("{\"req\": \"market.%s.trade.detail\",\"id\": \"%s\" }", symbol, clientId)
	return p.WebSocketClientBase.Send(req)
}

func (p *TradeWebSocketClient) Subscribe(symbol string, clientId string) error {
	sub := fmt.Sprintf("{\"sub\": \"market.%s.trade.detail\",\"id\": \"%s\" }", symbol, clientId)
	return p.WebSocketClientBase.Send(sub)
}

func (p *TradeWebSocketClient) UnSubscribe(symbol string, clientId string) error {
	unsub := fmt.Sprintf("{\"unsub\": \"market.%s.trade.detail\",\"id\": \"%s\" }", symbol, clientId)
	return p.WebSocketClientBase.Send(unsub)
}

func (p *TradeWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeTradeResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
