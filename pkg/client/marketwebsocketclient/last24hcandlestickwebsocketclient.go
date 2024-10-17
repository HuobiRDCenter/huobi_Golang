package marketwebsocketclient

import (
	"encoding/json"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client/websocketclientbase"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
)

// Responsible to handle last 24h candlestick data from WebSocket
type Last24hCandlestickWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

// Initializer
func (p *Last24hCandlestickWebSocketClient) Init(host string) *Last24hCandlestickWebSocketClient {
	p.WebSocketClientBase.Init(host)
	return p
}

// Set callback handler
func (p *Last24hCandlestickWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Request full candlestick data
func (p *Last24hCandlestickWebSocketClient) Request(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.detail", symbol)
	req := fmt.Sprintf("{\"req\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(req)

	applogger.Info("WebSocket requested, topic=%s, clientId=%s", topic, clientId)
}

// Subscribe latest 24h market stats
func (p *Last24hCandlestickWebSocketClient) Subscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.detail", symbol)
	sub := fmt.Sprintf("{\"sub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe latest 24 market stats
func (p *Last24hCandlestickWebSocketClient) UnSubscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.detail", symbol)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\",\"id\": \"%s\" }", symbol, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *Last24hCandlestickWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeLast24hCandlestickResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
