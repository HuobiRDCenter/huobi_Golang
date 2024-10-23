package accountwebsocketclient

import (
	"encoding/json"
	"fmt"

	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client/websocketclientbase"
	"github.com/huobirdcenter/huobi_golang/pkg/model/account"
)

// Responsible to handle account asset request from WebSocket
// This need authentication version 2
type SubscribeAccountWebSocketV2Client struct {
	websocketclientbase.WebSocketV2ClientBase
}

// Initializer
func (p *SubscribeAccountWebSocketV2Client) Init(accessKey string, secretKey string, host string, sign string) *SubscribeAccountWebSocketV2Client {
	p.WebSocketV2ClientBase.Init(accessKey, secretKey, host, sign)
	return p
}

// Set callback handler
func (p *SubscribeAccountWebSocketV2Client) SetHandler(
	authHandler websocketclientbase.AuthenticationV2ResponseHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketV2ClientBase.SetHandler(authHandler, p.handleMessage, responseHandler)
}

// Subscribe all balance updates of the current account
// 0: Only update when account balance changed
// 1: Update when either account balance changed or available balance changed
func (p *SubscribeAccountWebSocketV2Client) Subscribe(mode string, clientId string) {
	channel := fmt.Sprintf("accounts.update#%s", mode)
	sub := fmt.Sprintf("{\"action\":\"sub\", \"ch\":\"%s\", \"cid\": \"%s\"}", channel, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, channel=%s, clientId=%s", channel, clientId)
}

// Unsubscribe balance updates
func (p *SubscribeAccountWebSocketV2Client) UnSubscribe(mode string, clientId string) {
	channel := fmt.Sprintf("accounts.update#%s", mode)
	unsub := fmt.Sprintf("{\"action\":\"unsub\", \"ch\":\"%s\", \"cid\": \"%s\"}", channel, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, channel=%s, clientId=%s", channel, clientId)
}

func (p *SubscribeAccountWebSocketV2Client) handleMessage(msg string) (interface{}, error) {
	result := account.SubscribeAccountV2Response{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
