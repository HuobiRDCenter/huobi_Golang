package accountwebsocketclient

import (
	"../../response/account"
	"../websocketclientbase"
	"encoding/json"
	"fmt"
)

type SubscribeAccountWebSocketV1Client struct {
	websocketclientbase.WebSocketV1ClientBase
}

func (p *SubscribeAccountWebSocketV1Client) Init(accessKey string, secretKey string, host string) *SubscribeAccountWebSocketV1Client {
	p.WebSocketV1ClientBase.Init(accessKey, secretKey, host)
	return p
}

func (p *SubscribeAccountWebSocketV1Client) SetHandler(
	authHandler websocketclientbase.AuthenticationV1ResponseHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketV1ClientBase.SetHandler(authHandler, p.handleMessage, responseHandler)
}

func (p *SubscribeAccountWebSocketV1Client) Request(clientId string) error {

	req := fmt.Sprintf("{ \"op\":\"req\", \"topic\":\"accounts.list\", \"cid\": \"%s\"}", clientId)
	return p.Send(req)
}

func (p *SubscribeAccountWebSocketV1Client) Subscribe(mode string, clientId string) error {

	sub := fmt.Sprintf("{ \"op\":\"sub\", \"topic\":\"accounts\", \"mode\": \"%s\", \"cid\": \"%s\"}", mode, clientId)
	return p.Send(sub)
}

func (p *SubscribeAccountWebSocketV1Client) UnSubscribe(mode string, clientId string) error {
	unsub := fmt.Sprintf("{ \"op\":\"unsub\", \"topic\":\"accounts\", \"mode\": \"%s\", \"cid\": \"%s\" }", mode, clientId)
	return p.Send(unsub)
}
func (p *SubscribeAccountWebSocketV1Client) handleMessage(msg string) (interface{}, error) {
	result := account.SubscribeAccountV1Response{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}