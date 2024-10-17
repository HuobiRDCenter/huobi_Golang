package marketwebsocketclient

import (
	"encoding/json"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client/websocketclientbase"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
)

// Responsible to handle ticker data from WebSocket
type TickerWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

// Initializer
func (p *TickerWebSocketClient) Init(host string) *TickerWebSocketClient {
	p.WebSocketClientBase.Init(host)
	return p
}

// Set callback handler
func (p *TickerWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Request ticker data
func (p *TickerWebSocketClient) Request(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.ticker", symbol)
	req := fmt.Sprintf("{\"req\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(req)

	applogger.Info("WebSocket requested, topic=%s, clientId=%s", topic, clientId)
}

// Subscribe ticker market stats
func (p *TickerWebSocketClient) Subscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.ticker", symbol)
	sub := fmt.Sprintf("{\"sub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe ticker market stats
func (p *TickerWebSocketClient) UnSubscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.ticker", symbol)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\",\"id\": \"%s\" }", symbol, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *TickerWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeTickerResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
