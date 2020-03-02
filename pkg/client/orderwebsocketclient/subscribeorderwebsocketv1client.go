package orderwebsocketclient

import (
	"../../response/order"
	"../websocketclientbase"
	"encoding/json"
	"fmt"
)

type SubscribeOrderWebSocketV1Client struct {
	websocketclientbase.WebSocketV1ClientBase
}

func (p *SubscribeOrderWebSocketV1Client) Init(accessKey string, secretKey string, host string) *SubscribeOrderWebSocketV1Client {
	p.WebSocketV1ClientBase.Init(accessKey, secretKey, host)
	return p
}

func (p *SubscribeOrderWebSocketV1Client) SetHandler(
	authHandler websocketclientbase.AuthenticationV1ResponseHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketV1ClientBase.SetHandler(authHandler, p.handleMessage, responseHandler)
}

func (p *SubscribeOrderWebSocketV1Client) Subscribe(symbol string, clientId string) error {

	sub := fmt.Sprintf("{ \"op\":\"sub\", \"topic\":\"orders.%s.update\", \"cid\": \"%s\"}", symbol, clientId)
	return p.Send(sub)
}

func (p *SubscribeOrderWebSocketV1Client) UnSubscribe(symbol string, clientId string) error {
	unsub := fmt.Sprintf("{ \"op\":\"unsub\", \"topic\":\"orders.%s.update\", \"cid\": \"%s\"}", symbol, clientId)
	return p.Send(unsub)
}
func (p *SubscribeOrderWebSocketV1Client) handleMessage(msg string) (interface{}, error) {
	result := order.SubscribeOrderV1Response{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
