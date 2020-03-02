package marketwebsocketclient

import (
	"../../response/market"
	"../websocketclientbase"
	"encoding/json"
	"fmt"
)

type BestBidOfferWebSocketClient struct {
	websocketclientbase.WebSocketClientBase
}

func (p *BestBidOfferWebSocketClient) Init(host string) *BestBidOfferWebSocketClient {
	p.WebSocketClientBase.Init(host)
	return p
}

func (p *BestBidOfferWebSocketClient) SetHandler(
	connectedHandler websocketclientbase.ConnectedHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

func (p *BestBidOfferWebSocketClient) Subscribe(symbol string, clientId string) error {
	sub := fmt.Sprintf("{\"sub\": \"market.%s.bbo\", \"id\": \"%s\"}", symbol, clientId)
	return p.WebSocketClientBase.Send(sub)
}

func (p *BestBidOfferWebSocketClient) UnSubscribe(symbol string, clientId string) error {
	unsub := fmt.Sprintf("{\"unsub\": \"market.%s.bbo\", \"id\": \"%s\" }", symbol, clientId)
	return p.WebSocketClientBase.Send(unsub)
}

func (p *BestBidOfferWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubscribeBestBidOfferResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
