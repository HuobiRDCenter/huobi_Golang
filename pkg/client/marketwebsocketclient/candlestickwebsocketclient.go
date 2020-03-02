package marketwebsocketclient

import (
	"../../response/market"
	"../websocketclientbase"
	"encoding/json"
	"fmt"
)

type CandlestickWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

func (p *CandlestickWebSocketClient) Init(host string) *CandlestickWebSocketClient {
	p.WebSocketClientBase.Init(host)
	return p
}

func (p *CandlestickWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

func (p *CandlestickWebSocketClient) Request(symbol string, period string, from int64, to int64, clientId string) error {
	req := fmt.Sprintf("{\"req\": \"market.%s.kline.%s\", \"from\":%d, \"to\":%d, \"id\": \"%s\" }", symbol, period, from, to, clientId)
	return p.WebSocketClientBase.Send(req)
}

func (p *CandlestickWebSocketClient) Subscribe(symbol string, period string, clientId string) error {
	sub := fmt.Sprintf("{\"sub\": \"market.%s.kline.%s\", \"id\": \"%s\"}", symbol, period, clientId)
	return p.WebSocketClientBase.Send(sub)
}

func (p *CandlestickWebSocketClient) UnSubscribe(symbol string, period string, clientId string) error {
	unsub := fmt.Sprintf("{\"unsub\": \"market.%s.kline.%s\", \"id\": \"%s\" }", symbol, period, clientId)
	return p.WebSocketClientBase.Send(unsub)
}

func (p *CandlestickWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeCandlestickResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
