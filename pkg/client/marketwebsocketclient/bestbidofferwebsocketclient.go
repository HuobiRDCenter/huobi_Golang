package marketwebsocketclient

import (
	"encoding/json"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/pkg/client/websocketclientbase"
	"github.com/huobirdcenter/huobi_golang/pkg/response/market"
)

// Responsible to handle BBO data from WebSocket
type BestBidOfferWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

// Initializer
func (p *BestBidOfferWebSocketClient) Init(host string) *BestBidOfferWebSocketClient {
	p.WebSocketClientBase.Init(host)
	return p
}

// Set callback handler
func (p *BestBidOfferWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Subscribe latest market by price order book in snapshot mode at 1-second interval.
func (p *BestBidOfferWebSocketClient) Subscribe(symbol string, clientId string) error {
	sub := fmt.Sprintf("{\"sub\": \"market.%s.bbo\", \"id\": \"%s\"}", symbol, clientId)
	return p.WebSocketClientBase.Send(sub)
}

// Unsubscribe market by price order book
func (p *BestBidOfferWebSocketClient) UnSubscribe(symbol string, clientId string) error {
	unsub := fmt.Sprintf("{\"unsub\": \"market.%s.bbo\", \"id\": \"%s\" }", symbol, clientId)
	return p.WebSocketClientBase.Send(unsub)
}

func (p *BestBidOfferWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeBestBidOfferResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
