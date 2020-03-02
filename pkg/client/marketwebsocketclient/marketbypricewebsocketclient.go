package marketwebsocketclient

import (
	"../../response/market"
	"../websocketclientbase"
	"encoding/json"
	"fmt"
)

type MarketByPriceWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

func (p *MarketByPriceWebSocketClient) Init(host string) *MarketByPriceWebSocketClient {
	p.WebSocketClientBase.Init(host)
	return p
}

func (p *MarketByPriceWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

func (p *MarketByPriceWebSocketClient) Request(symbol string, clientId string) error {
	req := fmt.Sprintf("{\"req\": \"market.%s.mbp.150\",\"id\": \"%s\" }", symbol, clientId)
	return p.WebSocketClientBase.Send(req)
}

func (p *MarketByPriceWebSocketClient) Subscribe(symbol string, clientId string) error {
	sub := fmt.Sprintf("{\"sub\": \"market.%s.mbp.150\",\"id\": \"%s\" }", symbol, clientId)
	return p.WebSocketClientBase.Send(sub)
}

func (p *MarketByPriceWebSocketClient) UnSubscribe(symbol string, clientId string) error {
	unsub := fmt.Sprintf("{\"unsub\": \"market.%s.mbp.150\",\"id\": \"%s\" }", symbol, clientId)
	return p.WebSocketClientBase.Send(unsub)
}

func (p *MarketByPriceWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeMarketByPriceResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}