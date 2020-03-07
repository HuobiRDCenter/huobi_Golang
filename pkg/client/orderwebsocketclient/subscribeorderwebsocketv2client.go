package orderwebsocketclient

import (
	"../../response/order"
	"../websocketclientbase"
	"encoding/json"
	"fmt"
)

// Responsible to handle trade clear from WebSocket
// This need authentication version 2
type SubscribeOrderWebSocketV2Client struct {
	websocketclientbase.WebSocketV2ClientBase
}

// Initializer
func (p *SubscribeOrderWebSocketV2Client) Init(accessKey string, secretKey string, host string) *SubscribeOrderWebSocketV2Client {
	p.WebSocketV2ClientBase.Init(accessKey, secretKey, host)
	return p
}

// Set callback handler
func (p *SubscribeOrderWebSocketV2Client) SetHandler(
	authHandler websocketclientbase.AuthenticationV2ResponseHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketV2ClientBase.SetHandler(authHandler, p.handleMessage, responseHandler)
}

// Subscribe trade details including transaction fee and transaction fee deduction etc.
// It only updates when transaction occurs.
func (p *SubscribeOrderWebSocketV2Client) Subscribe(mode string, clientId string) error {
	sub := fmt.Sprintf("{\"action\":\"sub\", \"ch\":\"accounts.update#%s\", \"cid\": \"%s\"}", mode, clientId)
	return p.Send(sub)
}

// Unsubscribe trade update
func (p *SubscribeOrderWebSocketV2Client) UnSubscribe(mode string, clientId string) error {
	unsub := fmt.Sprintf("{\"action\":\"unsub\", \"ch\":\"accounts.update#%s\", \"cid\": \"%s\"}", mode, clientId)
	return p.Send(unsub)
}
func (p *SubscribeOrderWebSocketV2Client) handleMessage(msg string) (interface{}, error) {
	result := &order.SubscribeOrderV2Response{}
	err := json.Unmarshal([]byte(msg), result)
	return result, err
}