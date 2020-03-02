package orderwebsocketclient

import (
	"../../getrequest"
	"../../postrequest"
	"../../response/order"
	"../websocketclientbase"
	"encoding/json"
)

type RequestOrdersWebSocketV1Client struct {
	websocketclientbase.WebSocketV1ClientBase
}

func (p *RequestOrdersWebSocketV1Client) Init(accessKey string, secretKey string, host string) *RequestOrdersWebSocketV1Client {
	p.WebSocketV1ClientBase.Init(accessKey, secretKey, host)
	return p
}

func (p *RequestOrdersWebSocketV1Client) SetHandler(
	authHandler websocketclientbase.AuthenticationV1ResponseHandler,
	responseHandler websocketclientbase.ResponseHandler) {
	p.WebSocketV1ClientBase.SetHandler(authHandler, p.handleMessage, responseHandler)
}

func (p *RequestOrdersWebSocketV1Client) Request(req getrequest.RequestOrdersRequest) error {

	reqString, _ := postrequest.ToJson(req)
	return p.Send(reqString)
}

func (p *RequestOrdersWebSocketV1Client) handleMessage(msg string) (interface{}, error) {
	result := order.RequestOrdersV1Response{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
