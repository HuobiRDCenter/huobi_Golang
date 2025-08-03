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
	subOrderUpdateV2()
	subTradeClear()
}

func subOrderUpdateV2() {
	// Initialize a new instance
	client := new(orderwebsocketclient.SubscribeOrderWebSocketV2Client).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)

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
							o.Event, o.Symbol, o.Type, o.OrderStatus)
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
	client := new(orderwebsocketclient.SubscribeTradeClearWebSocketV2Client).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)

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
