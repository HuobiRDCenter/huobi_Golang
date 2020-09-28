package accountwebsocketclientexample

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client/accountwebsocketclient"
	"github.com/huobirdcenter/huobi_golang/pkg/model/account"
	"github.com/huobirdcenter/huobi_golang/pkg/model/auth"
)

func RunAllExamples() {
	reqAccountUpdateV1()
	subAccountUpdateV1()
	subAccountUpdateV2()
}

func reqAccountUpdateV1() {
	// Initialize a new instance
	client := new(accountwebsocketclient.RequestAccountWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)

	// Set the callback handlers
	client.SetHandler(
		// Authentication response handler
		func(resp *auth.WebSocketV1AuthenticationResponse) {
			if resp.ErrorCode == 0 {
				err := client.Request("")
				if err != nil {
					applogger.Error("Request error: %s", err)
				} else {
					applogger.Debug("Sent request")
				}
			} else {
				applogger.Error("Authentication error: %d", resp.ErrorCode)
			}

		},
		// Response handler
		func(resp interface{}) {
			reqResponse, ok := resp.(account.RequestAccountV1Response)
			if ok {
				if &reqResponse.Data != nil {
					for _, a := range reqResponse.Data {
						applogger.Info("Account id: %d, type: %s, state: %s", a.Id, a.Type, a.State)
						if &a.List != nil {
							for _, b := range a.List {
								applogger.Info("Currency: %s, type: %s, balance: %s", b.Currency, b.Type, b.Balance)
							}
						}
					}
				}
			} else {
				applogger.Warn("Received unknown response: %+v", resp)
			}
		})

	// Connect to the server and wait for the handler to handle the response
	err := client.Connect(false)
	if err != nil {
		applogger.Error("Client Connect error: %s", err)
		return
	}

	// Unsubscribe the topic
	fmt.Println("Press ENTER to stop...")
	fmt.Scanln()

	// Close the connection
	client.Close()
	applogger.Info("Client closed")
}

func subAccountUpdateV1() {
	// Initialize a new instance
	client := new(accountwebsocketclient.SubscribeAccountWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)

	// Set the callback handlers
	client.SetHandler(
		// Authentication response handler
		func(resp *auth.WebSocketV1AuthenticationResponse) {
			if resp.ErrorCode == 0 {
				err := client.Subscribe("1", "1250")
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
			subResponse, ok := resp.(account.SubscribeAccountV1Response)
			if ok {
				if &subResponse.Data != nil {
					applogger.Info("Account update event: %s", subResponse.Data.Event)
					if &subResponse.Data.List != nil {
						for _, b := range subResponse.Data.List {
							applogger.Info("Account id: %d, currency: %s, type: %s, balance: %s", b.AccountId, b.Currency, b.Type, b.Balance)
						}
					}
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

	// Unsubscribe the topic
	err = client.UnSubscribe("1", "1250")
	if err != nil {
		applogger.Error("UnSubscribe error: %s", err)
	}

	// Close the connection
	client.Close()
	applogger.Info("Client closed")
}

func subAccountUpdateV2() {
	// Initialize a new instance
	client := new(accountwebsocketclient.SubscribeAccountWebSocketV2Client).Init(config.AccessKey, config.SecretKey, config.Host)

	// Set the callback handlers
	client.SetHandler(
		// Authentication response handler
		func(resp *auth.WebSocketV2AuthenticationResponse) {
			if resp.IsSuccess() {
				client.Subscribe("1", "1149")
			} else {
				applogger.Error("Authentication error, code: %d, message:%s", resp.Code, resp.Message)
			}
		},
		// Response handler
		func(resp interface{}) {
			subResponse, ok := resp.(account.SubscribeAccountV2Response)
			if ok {
				if subResponse.Action == "sub" {
					if subResponse.IsSuccess() {
						applogger.Info("Subscription topic %s successfully", subResponse.Ch)
					} else {
						applogger.Error("Subscription topic %s error, code: %d, message: %s", subResponse.Ch, subResponse.Code, subResponse.Message)
					}
				} else if subResponse.Action == "push" {
					if subResponse.Data != nil {
						b := subResponse.Data
						if b.ChangeTime == 0 {
							applogger.Info("Account overview, id: %d, currency: %s, balance: %s", b.AccountId, b.Currency, b.Balance)
						} else {
							applogger.Info("Account update, id: %d, currency: %s, balance: %s, time: %d", b.AccountId, b.Currency, b.Balance, b.ChangeTime)
						}
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

	// Unsubscribe the topic
	client.UnSubscribe("1", "1250")

	// Close the connection
	client.Close()
	applogger.Info("Client closed")
}
