package marketwebsocketclient

import (
	"encoding/json"
	"fmt"
	"github.com/newgoo/huobi_golang/logging/applogger"
	"github.com/newgoo/huobi_golang/pkg/client/websocketclientbase"
	"github.com/newgoo/huobi_golang/pkg/model/market"
)

// Responsible to handle Trade data from WebSocket
type TradeWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

// Initializer
func (p *TradeWebSocketClient) Init(host string) *TradeWebSocketClient {
	p.WebSocketClientBase.Init(host)
	return p
}

// Set callback handler
func (p *TradeWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Request latest 300 trade data
func (p *TradeWebSocketClient) Request(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.trade.detail", symbol)
	req := fmt.Sprintf("{\"req\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(req)

	applogger.Info("WebSocket requested, topic=%s, clientId=%s", topic, clientId)
}

// Subscribe latest completed trade in tick by tick mode
func (p *TradeWebSocketClient) Subscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.trade.detail", symbol)
	sub := fmt.Sprintf("{\"sub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe trade
func (p *TradeWebSocketClient) UnSubscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.trade.detail", symbol)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *TradeWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeTradeResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
