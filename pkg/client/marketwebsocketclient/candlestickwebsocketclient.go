package marketwebsocketclient

import (
	"../../response/market"
	"../websocketclientbase"
	"encoding/json"
	"fmt"
)

// Responsible to handle candlestick data from WebSocket
type CandlestickWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

// Initializer
func (p *CandlestickWebSocketClient) Init(host string) *CandlestickWebSocketClient {
	p.WebSocketClientBase.Init(host)
	return p
}

// Set callback handler
func (p *CandlestickWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Request the full candlestick data according to specified criteria
func (p *CandlestickWebSocketClient) Request(symbol string, period string, from int64, to int64, clientId string) error {
	req := fmt.Sprintf("{\"req\": \"market.%s.kline.%s\", \"from\":%d, \"to\":%d, \"id\": \"%s\" }", symbol, period, from, to, clientId)
	return p.WebSocketClientBase.Send(req)
}

// Subscribe candlestick data
func (p *CandlestickWebSocketClient) Subscribe(symbol string, period string, clientId string) error {
	sub := fmt.Sprintf("{\"sub\": \"market.%s.kline.%s\", \"id\": \"%s\"}", symbol, period, clientId)
	return p.WebSocketClientBase.Send(sub)
}

// Unsubscribe candlestick data
func (p *CandlestickWebSocketClient) UnSubscribe(symbol string, period string, clientId string) error {
	unsub := fmt.Sprintf("{\"unsub\": \"market.%s.kline.%s\", \"id\": \"%s\" }", symbol, period, clientId)
	return p.WebSocketClientBase.Send(unsub)
}

func (p *CandlestickWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeCandlestickResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
