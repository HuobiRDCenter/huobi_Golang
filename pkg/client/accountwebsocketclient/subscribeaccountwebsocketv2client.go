package accountwebsocketclient

import (
	"../../response/account"
	"../websocketclientbase"
	"encoding/json"
	"fmt"
)

type SubscribeAccountWebSocketV2Client struct {
	websocketclientbase.WebSocketV2ClientBase
}

func (p *SubscribeAccountWebSocketV2Client) Init(accessKey string, secretKey string, host string) *SubscribeAccountWebSocketV2Client {
	p.WebSocketV2ClientBase.Init(accessKey, secretKey, host)
	return p
}

func (p *SubscribeAccountWebSocketV2Client) SetHandler(
	authHandler websocketclientbase.AuthenticationV2ResponseHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketV2ClientBase.SetHandler(authHandler, p.handleMessage, responseHandler)
}

func (p *SubscribeAccountWebSocketV2Client) Subscribe(mode string, clientId string) error {
	sub := fmt.Sprintf("{\"action\":\"sub\", \"ch\":\"accounts.update#%s\", \"cid\": \"%s\"}", mode, clientId)
	return p.Send(sub)
}

func (p *SubscribeAccountWebSocketV2Client) UnSubscribe(mode string, clientId string) error {
	unsub := fmt.Sprintf("{\"action\":\"unsub\", \"ch\":\"accounts.update#%s\", \"cid\": \"%s\"}", mode, clientId)
	return p.Send(unsub)
}
func (p *SubscribeAccountWebSocketV2Client) handleMessage(msg string) (interface{}, error) {
	result := &account.SubscribeAccountV2Response{}
	err := json.Unmarshal([]byte(msg), result)
	return result, err
}
