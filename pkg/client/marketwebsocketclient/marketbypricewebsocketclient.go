package marketwebsocketclient

import (
	"encoding/json"
	"fmt"
	"github.com/newgoo/huobi_golang/logging/applogger"
	"github.com/newgoo/huobi_golang/pkg/client/websocketclientbase"
	"github.com/newgoo/huobi_golang/pkg/model/market"
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
func (p *MarketByPriceWebSocketClient) Request(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.mbp.150", symbol)
	req := fmt.Sprintf("{\"req\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.WebSocketClientBase.Send(req)

	applogger.Info("WebSocket requested, topic=%s, clientId=%s", topic, clientId)
}

// Subscribe incremental update of Market By Price order book
func (p *MarketByPriceWebSocketClient) Subscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.mbp.150", symbol)
	sub := fmt.Sprintf("{\"sub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.WebSocketClientBase.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Subscribe full Market By Price order book
func (p *MarketByPriceWebSocketClient) SubscribeFull(symbol string, level int, clientId string) {
	topic := fmt.Sprintf("market.%s.mbp.refresh.%d", symbol, level)
	sub := fmt.Sprintf("{\"sub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe update of Market By Price order book
func (p *MarketByPriceWebSocketClient) UnSubscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.mbp.150", symbol)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe full Market By Price order book
func (p *MarketByPriceWebSocketClient) UnSubscribeFull(symbol string, level int, clientId string) {
	topic := fmt.Sprintf("market.%s.mbp.refresh.%d", symbol, level)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *MarketByPriceWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeMarketByPriceResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
