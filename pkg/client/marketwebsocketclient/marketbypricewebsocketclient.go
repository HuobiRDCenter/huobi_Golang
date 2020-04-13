package marketwebsocketclient

import (
	"encoding/json"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/pkg/client/websocketclientbase"
	"github.com/huobirdcenter/huobi_golang/pkg/response/market"
)

// Responsible to handle MBP data from WebSocket
type MarketByPriceWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

// Initializer
func (p *MarketByPriceWebSocketClient) Init(host string) *MarketByPriceWebSocketClient {
	p.WebSocketClientBase.Init(host)
	return p
}

// Set callback handler
func (p *MarketByPriceWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Request full Market By Price order book
func (p *MarketByPriceWebSocketClient) Request(symbol string, clientId string) error {
	req := fmt.Sprintf("{\"req\": \"market.%s.mbp.150\",\"id\": \"%s\" }", symbol, clientId)
	return p.WebSocketClientBase.Send(req)
}

// Subscribe incremental update of Market By Price order book
func (p *MarketByPriceWebSocketClient) Subscribe(symbol string, clientId string) error {
	sub := fmt.Sprintf("{\"sub\": \"market.%s.mbp.150\",\"id\": \"%s\" }", symbol, clientId)
	return p.WebSocketClientBase.Send(sub)
}

// Subscribe full Market By Price order book
func (p *MarketByPriceWebSocketClient) SubscribeFull(symbol string, level int, clientId string) error {
	sub := fmt.Sprintf("{\"sub\": \"market.%s.mbp.refresh.%d\",\"id\": \"%s\" }", symbol, level, clientId)
	return p.WebSocketClientBase.Send(sub)
}

// Unsubscribe update of Market By Price order book
func (p *MarketByPriceWebSocketClient) UnSubscribe(symbol string, clientId string) error {
	unsub := fmt.Sprintf("{\"unsub\": \"market.%s.mbp.150\",\"id\": \"%s\" }", symbol, clientId)
	return p.WebSocketClientBase.Send(unsub)
}

// Unsubscribe full Market By Price order book
func (p *MarketByPriceWebSocketClient) UnSubscribeFull(symbol string, level int, clientId string) error {
	unsub := fmt.Sprintf("{\"unsub\": \"market.%s.mbp.refresh.%d\",\"id\": \"%s\" }", symbol, level, clientId)
	return p.WebSocketClientBase.Send(unsub)
}

func (p *MarketByPriceWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeMarketByPriceResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
