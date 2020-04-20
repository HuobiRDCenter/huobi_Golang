package orderwebsocketclient

import (
	"encoding/json"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/pkg/client/websocketclientbase"
	"github.com/huobirdcenter/huobi_golang/pkg/response/order"
)

// Responsible to handle trade clear from WebSocket
// This need authentication version 2
type SubscribeTradeClearWebSocketV2Client struct {
	websocketclientbase.WebSocketV2ClientBase
}

// Initializer
func (p *SubscribeTradeClearWebSocketV2Client) Init(accessKey string, secretKey string, host string) *SubscribeTradeClearWebSocketV2Client {
	p.WebSocketV2ClientBase.Init(accessKey, secretKey, host)
	return p
}

// Set callback handler
func (p *SubscribeTradeClearWebSocketV2Client) SetHandler(
	authHandler websocketclientbase.AuthenticationV2ResponseHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketV2ClientBase.SetHandler(authHandler, p.handleMessage, responseHandler)
}

// Subscribe trade details including transaction fee and transaction fee deduction etc.
// It only updates when transaction occurs.
func (p *SubscribeTradeClearWebSocketV2Client) Subscribe(symbol string, clientId string) error {
	sub := fmt.Sprintf("{\"action\":\"sub\", \"ch\":\"trade.clearing#%s\", \"cid\": \"%s\"}", symbol, clientId)
	return p.Send(sub)
}

// Unsubscribe trade update
func (p *SubscribeTradeClearWebSocketV2Client) UnSubscribe(symbol string, clientId string) error {
	unsub := fmt.Sprintf("{\"action\":\"unsub\", \"ch\":\"trade.clearing#%s\", \"cid\": \"%s\"}", symbol, clientId)
	return p.Send(unsub)
}
func (p *SubscribeTradeClearWebSocketV2Client) handleMessage(msg string) (interface{}, error) {
	result := order.SubscribeTradeClearResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
