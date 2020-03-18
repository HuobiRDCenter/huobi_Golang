package orderwebsocketclientexample

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/internal/model"
	"github.com/huobirdcenter/huobi_golang/pkg/client/orderwebsocketclient"
	"github.com/huobirdcenter/huobi_golang/pkg/getrequest"
	"github.com/huobirdcenter/huobi_golang/pkg/response/order"
)

func RunAllExamples() {
	reqOrderV1()
	reqOrdersV1()
	subOrderUpdateV1()
	subOrderUpdateV2()
}

func reqOrderV1() {
	client := new(orderwebsocketclient.RequestOrderWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)
	client.SetHandler(
		func(resp *model.WebSocketV1AuthenticationResponse) {
			if resp.ErrorCode == 0 {
				err := client.Request("1", "1601")
				if err != nil {
					fmt.Printf("Request error: %s\n", err)
				} else {
					fmt.Println("Sent request")
				}
			} else {
				fmt.Printf("Authentication error: %d\n", resp.ErrorCode)
			}

		},
		func(resp interface{}) {
			reqResponse, ok := resp.(order.RequestOrderV1Response)
			if ok {
				if &reqResponse.Data != nil {
					o := reqResponse.Data
					fmt.Printf("Request order, id: %d, state: %s, symbol: %s, price: %s, filled amount: %s", o.Id, o.State, o.Symbol, o.Price, o.FilledAmount)
				} else {
					fmt.Printf("Request order error: %s", reqResponse.ErrorCode)
				}
			} else {
				fmt.Printf("Received unknown response: %v\n", resp)
			}
		})

	err := client.Connect(true)
	if err != nil {
		fmt.Printf("Client Connect error: %s\n", err)
		return
	}

	fmt.Println("Press ENTER to stop...")
	fmt.Scanln()

	client.Close()
	fmt.Println("Client closed")
}

func reqOrdersV1() {
	client := new(orderwebsocketclient.RequestOrdersWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)
	client.SetHandler(
		func(resp *model.WebSocketV1AuthenticationResponse) {
			if resp.ErrorCode == 0 {
				req := getrequest.RequestOrdersRequest{
					AccountId: 11136102,
					Symbol:    "btcusdt",
					States:    "submitted, created, canceled",
				}
				err := client.Request(req)
				if err != nil {
					fmt.Printf("Request error: %s\n", err)
				} else {
					fmt.Println("Sent request")
				}
			} else {
				fmt.Printf("Authentication error: %d\n", resp.ErrorCode)
			}

		},
		func(resp interface{}) {
			reqResponse, ok := resp.(order.RequestOrdersV1Response)
			if ok {
				if &reqResponse.Data != nil {
					for _, o := range reqResponse.Data {
						fmt.Printf("Request order, id: %d, state: %s, symbol: %s, price: %s, filled amount: %s", o.Id, o.State, o.Symbol, o.Price, o.FilledAmount)
					}
				} else {
					fmt.Printf("Request order error: %s", reqResponse.ErrorCode)
				}
			} else {
				fmt.Printf("Received unknown response: %v\n", resp)
			}
		})

	err := client.Connect(true)
	if err != nil {
		fmt.Printf("Client Connect error: %s\n", err)
		return
	}

	fmt.Println("Press ENTER to stop...")
	fmt.Scanln()

	client.Close()
	fmt.Println("Client closed")
}

func subOrderUpdateV1() {
	// Initialize a new instance
	client := new(orderwebsocketclient.SubscribeOrderWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)

	// Set the callback handlers
	client.SetHandler(
		// Connected handler
		func(resp *model.WebSocketV1AuthenticationResponse) {
			if resp.ErrorCode == 0 {
				err := client.Subscribe("btcusdt", "1601")
				if err != nil {
					fmt.Printf("Subscribe error: %s\n", err)
				} else {
					fmt.Println("Sent subscription")
				}
			} else {
				fmt.Printf("Authentication error: %d\n", resp.ErrorCode)
			}

		},
		// Response handler
		func(resp interface{}) {
			subResponse, ok := resp.(order.SubscribeOrderV1Response)
			if ok {
				if &subResponse.Data != nil {
					o := subResponse.Data
					fmt.Printf("Order update, id: %d, state: %s, symbol: %s, price: %s, filled amount: %s", o.OrderId, o.OrderState, o.Symbol, o.Price, o.FilledAmount)
				}
			} else {
				fmt.Printf("Received unknown response: %v\n", resp)
			}
		})

	// Connect to the server and wait for the handler to handle the response
	err := client.Connect(true)
	if err != nil {
		fmt.Printf("Client Connect error: %s\n", err)
		return
	}

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	err = client.UnSubscribe("1", "1250")
	if err != nil {
		fmt.Printf("UnSubscribe error: %s\n", err)
	}

	client.Close()
	fmt.Println("Client closed")
}

func subOrderUpdateV2() {
	// Initialize a new instance
	client := new(orderwebsocketclient.SubscribeOrderWebSocketV2Client).Init(config.AccessKey, config.SecretKey, config.Host)

	// Set the callback handlers
	client.SetHandler(
		// Connected handler
		func(resp *model.WebSocketV2AuthenticationResponse) {
			if resp.IsAuth() {
				// Subscribe if authentication passed
				err := client.Subscribe("1", "1149")
				if err != nil {
					fmt.Printf("Subscribe error: %s\n", err)
				} else {
					fmt.Println("Sent subscription")
				}
			} else {
				fmt.Printf("Authentication error: %d\n", resp.Code)
			}
		},
		// Response handler
		func(resp interface{}) {
			subResponse, ok := resp.(order.SubscribeOrderV2Response)
			if ok {
				if &subResponse.Data != nil {
					o := subResponse.Data
					fmt.Printf("Order update, symbol: %s, order id: %d, price: %s, volume: %s",
						o.Symbol, o.OrderId, o.TradePrice, o.TradeVolume)
				}
			} else {
				fmt.Printf("Received unknown response: %v\n", resp)
			}
		})

	// Connect to the server and wait for the handler to handle the response
	err := client.Connect(true)
	if err != nil {
		fmt.Printf("Client Connect error: %s\n", err)
		return
	}

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	err = client.UnSubscribe("1", "1250")
	if err != nil {
		fmt.Printf("UnSubscribe error: %s\n", err)
	}

	client.Close()
	fmt.Println("Client closed")
}
