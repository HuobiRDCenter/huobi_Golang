package orderwebsocketclientexample

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client/orderwebsocketclient"
	"github.com/huobirdcenter/huobi_golang/pkg/model/auth"
	"github.com/huobirdcenter/huobi_golang/pkg/model/order"
)

func RunAllExamples() {
	reqOrderV1()
	reqOrdersV1()
	subOrderUpdateV1()
	subOrderUpdateV2()
	subTradeClear()
}

func reqOrderV1() {
	client := new(orderwebsocketclient.RequestOrderWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)
	client.SetHandler(
		func(resp *auth.WebSocketV1AuthenticationResponse) {
			if resp.ErrorCode == 0 {
				err := client.Request("1", "1601")
				if err != nil {
					applogger.Error("Request error: %s", err)
				} else {
					applogger.Debug("Sent request")
				}
			} else {
				applogger.Error("Authentication error: %d", resp.ErrorCode)
			}

		},
		func(resp interface{}) {
			reqResponse, ok := resp.(order.RequestOrderV1Response)
			if ok {
				if &reqResponse.Data != nil {
					o := reqResponse.Data
					applogger.Info("Request order, id: %d, state: %s, symbol: %s, price: %s, filled amount: %s", o.Id, o.State, o.Symbol, o.Price, o.FilledAmount)
				} else {
					applogger.Error("Request order error: %s", reqResponse.ErrorCode)
				}
			} else {
				applogger.Warn("Received unknown response: %v", resp)
			}
		})

	err := client.Connect(true)
	if err != nil {
		applogger.Error("Client Connect error: %s", err)
		return
	}

	fmt.Println("Press ENTER to stop...")
	fmt.Scanln()

	client.Close()
	applogger.Info("Client closed")
}

func reqOrdersV1() {
	client := new(orderwebsocketclient.RequestOrdersWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)
	client.SetHandler(
		func(resp *auth.WebSocketV1AuthenticationResponse) {
			if resp.ErrorCode == 0 {
				req := order.RequestOrdersRequest{
					AccountId: 11136102,
					Symbol:    "btcusdt",
					States:    "submitted, created, canceled",
				}
				err := client.Request(req)
				if err != nil {
					applogger.Error("Request error: %s", err)
				} else {
					applogger.Debug("Sent request")
				}
			} else {
				applogger.Error("Authentication error: %d", resp.ErrorCode)
			}

		},
		func(resp interface{}) {
			reqResponse, ok := resp.(order.RequestOrdersV1Response)
			if ok {
				if &reqResponse.Data != nil {
					for _, o := range reqResponse.Data {
						applogger.Info("Request order, id: %d, state: %s, symbol: %s, price: %s, filled amount: %s", o.Id, o.State, o.Symbol, o.Price, o.FilledAmount)
					}
				} else {
					applogger.Error("Request order error: %s", reqResponse.ErrorCode)
				}
			} else {
				applogger.Warn("Received unknown response: %+v", resp)
			}
		})

	err := client.Connect(true)
	if err != nil {
		applogger.Error("Client Connect error: %s", err)
		return
	}

	fmt.Println("Press ENTER to stop...")
	fmt.Scanln()

	client.Close()
	applogger.Info("Client closed")
}

func subOrderUpdateV1() {
	// Initialize a new instance
	client := new(orderwebsocketclient.SubscribeOrderWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)

	// Set the callback handlers
	client.SetHandler(
		// Connected handler
		func(resp *auth.WebSocketV1AuthenticationResponse) {
			if resp.ErrorCode == 0 {
				err := client.Subscribe("btcusdt", "1601")
				if err != nil {
					applogger.Error("Subscribe error: %s", err)
				} else {
					applogger.Debug("Sent subscription")
				}
			} else {
				applogger.Error("Authentication error: %d", resp.ErrorCode)
			}

		},
		// Response handler
		func(resp interface{}) {
			subResponse, ok := resp.(order.SubscribeOrderV1Response)
			if ok {
				if &subResponse.Data != nil {
					o := subResponse.Data
					applogger.Info("Order update, id: %d, state: %s, symbol: %s, price: %s, filled amount: %s", o.OrderId, o.OrderState, o.Symbol, o.Price, o.FilledAmount)
				}
			} else {
				applogger.Warn("Received unknown response: %+v", resp)
			}
		})

	// Connect to the server and wait for the handler to handle the response
	err := client.Connect(true)
	if err != nil {
		applogger.Error("Client Connect error: %s", err)
		return
	}

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	err = client.UnSubscribe("1", "1250")
	if err != nil {
		applogger.Error("UnSubscribe error: %s", err)
	}

	client.Close()
	applogger.Info("Client closed")
}

func subOrderUpdateV2() {
	// Initialize a new instance
	client := new(orderwebsocketclient.SubscribeOrderWebSocketV2Client).Init(config.AccessKey, config.SecretKey, config.Host)

	// Set the callback handlers
	client.SetHandler(
		// Connected handler
		func(resp *auth.WebSocketV2AuthenticationResponse) {
			if resp.IsSuccess() {
				// Subscribe if authentication passed
				client.Subscribe("btcusdt", "1149")
			} else {
				applogger.Error("Authentication error, code: %d, message:%s", resp.Code, resp.Message)
			}
		},
		// Response handler
		func(resp interface{}) {
			subResponse, ok := resp.(order.SubscribeOrderV2Response)
			if ok {
				if subResponse.Action == "sub" {
					if subResponse.IsSuccess() {
						applogger.Info("Subscription topic %s successfully", subResponse.Ch)
					} else {
						applogger.Error("Subscription topic %s error, code: %d, message: %s", subResponse.Ch, subResponse.Code, subResponse.Message)
					}
				} else if subResponse.Action == "push" {
					if subResponse.Data != nil {
						o := subResponse.Data
						applogger.Info("Order update, event: %s, symbol: %s, type: %s, status: %s",
							o.EventType, o.Symbol, o.Type, o.OrderStatus)
					}
				}
			} else {
				applogger.Warn("Received unknown response: %v", resp)
			}
		})

	// Connect to the server and wait for the handler to handle the response
	client.Connect(true)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	client.UnSubscribe("1", "1250")

	client.Close()
	applogger.Info("Client closed")
}

func subTradeClear() {
	// Initialize a new instance
	client := new(orderwebsocketclient.SubscribeTradeClearWebSocketV2Client).Init(config.AccessKey, config.SecretKey, config.Host)

	// Set the callback handlers
	client.SetHandler(
		// Connected handler
		func(resp *auth.WebSocketV2AuthenticationResponse) {
			if resp.IsSuccess() {
				// Subscribe if authentication passed
				client.Subscribe("btcusdt", "1149")
			} else {
				applogger.Error("Authentication error, code: %d, message:%s", resp.Code, resp.Message)
			}
		},
		// Response handler
		func(resp interface{}) {
			subResponse, ok := resp.(order.SubscribeTradeClearResponse)
			if ok {
				if subResponse.Action == "sub" {
					if subResponse.IsSuccess() {
						applogger.Info("Subscription topic %s successfully", subResponse.Ch)
					} else {
						applogger.Error("Subscription topic %s error, code: %d, message: %s", subResponse.Ch, subResponse.Code, subResponse.Message)
					}
				} else if subResponse.Action == "push" {
					if subResponse.Data != nil {
						o := subResponse.Data
						applogger.Info("Order update, symbol: %s, order id: %d, price: %s, volume: %s",
							o.Symbol, o.OrderId, o.TradePrice, o.TradeVolume)
					}
				}
			} else {
				applogger.Warn("Received unknown response: %v", resp)
			}
		})

	// Connect to the server and wait for the handler to handle the response
	client.Connect(true)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	client.UnSubscribe("btcusdt", "1250")

	client.Close()
	applogger.Info("Client closed")
}
