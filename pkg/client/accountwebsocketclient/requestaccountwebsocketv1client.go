package accountwebsocketclient

import (
	"../../response/account"
	"../websocketclientbase"
	"encoding/json"
	"fmt"
)

type RequestAccountWebSocketV1Client struct {
	websocketclientbase.WebSocketV1ClientBase
}

func (p *RequestAccountWebSocketV1Client) Init(accessKey string, secretKey string, host string) *RequestAccountWebSocketV1Client {
	p.WebSocketV1ClientBase.Init(accessKey, secretKey, host)
	return p
}

func (p *RequestAccountWebSocketV1Client) SetHandler(
	authHandler websocketclientbase.AuthenticationV1ResponseHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketV1ClientBase.SetHandler(authHandler, p.handleMessage, responseHandler)
}

func (p *RequestAccountWebSocketV1Client) Request(clientId string) error {

	req := fmt.Sprintf("{ \"op\":\"req\", \"topic\":\"accounts.list\", \"cid\": \"%s\"}", clientId)
	return p.Send(req)
}

func (p *RequestAccountWebSocketV1Client) Subscribe(mode string, clientId string) error {

	sub := fmt.Sprintf("{ \"op\":\"sub\", \"topic\":\"accounts\", \"mode\": \"%s\", \"cid\": \"%s\"}", mode, clientId)
	return p.Send(sub)
}

func (p *RequestAccountWebSocketV1Client) UnSubscribe(mode string, clientId string) error {
	unsub := fmt.Sprintf("{ \"op\":\"unsub\", \"topic\":\"accounts\", \"mode\": \"%s\", \"cid\": \"%s\" }", mode, clientId)
	return p.Send(unsub)
}
func (p *RequestAccountWebSocketV1Client) handleMessage(msg string) (interface{}, error) {
	result := account.RequestAccountV1Response{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}