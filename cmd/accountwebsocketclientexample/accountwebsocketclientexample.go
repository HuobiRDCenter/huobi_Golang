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
	subAccountUpdateV2()
}

func subAccountUpdateV2() {
	// Initialize a new instance
	client := new(accountwebsocketclient.SubscribeAccountWebSocketV2Client).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)

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
