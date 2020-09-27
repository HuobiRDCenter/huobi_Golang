package marketwebsocketclient

import (
	"encoding/json"
	"fmt"
	"github.com/newgoo/huobi_golang/logging/applogger"
	"github.com/newgoo/huobi_golang/pkg/client/websocketclientbase"
	"github.com/newgoo/huobi_golang/pkg/model/market"
)

// Responsible to handle MBP data from WebSocket
type MarketByPriceTickWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

// Initializer
func (p *MarketByPriceTickWebSocketClient) Init(host string) *MarketByPriceTickWebSocketClient {
	p.WebSocketClientBase.InitWithFeedPath(host)
	return p
}

// Set callback handler
func (p *MarketByPriceTickWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Request full Market By Price order book, level: 5, 20, 150
func (p *MarketByPriceTickWebSocketClient) Request(symbol string, level int, clientId string) {
	topic := fmt.Sprintf("market.%s.mbp.%d", symbol, level)
	req := fmt.Sprintf("{\"req\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.WebSocketClientBase.Send(req)

	applogger.Info("WebSocket requested, topic=%s, clientId=%s", topic, clientId)
}

// Subscribe incremental update of Market By Price order book, level: 5, 20, 150
func (p *MarketByPriceTickWebSocketClient) Subscribe(symbol string, level int, clientId string) {
	topic := fmt.Sprintf("market.%s.mbp.%d", symbol, level)
	sub := fmt.Sprintf("{\"sub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.WebSocketClientBase.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe update of Market By Price order book
func (p *MarketByPriceTickWebSocketClient) UnSubscribe(symbol string, level int, clientId string) {
	topic := fmt.Sprintf("market.%s.mbp.%d", symbol, level)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *MarketByPriceTickWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeMarketByPriceResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
