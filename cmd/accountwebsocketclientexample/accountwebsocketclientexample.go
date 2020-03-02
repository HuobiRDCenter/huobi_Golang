package accountwebsocketclientexample

import (
	"../../config"
	"../../internal/model"
	"../../pkg/client/accountwebsocketclient"
	"../../pkg/response/account"
	"fmt"
)

func RunAllExamples() {
	reqAccountUpdateV1()
	subAccountUpdateV1()
	subAccountUpdateV2()
}

func reqAccountUpdateV1() {
	client := new(accountwebsocketclient.RequestAccountWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)
	client.SetHandler(
		func(resp *model.WebSocketV1AuthenticationResponse) {
			if resp.ErrorCode == 0 {
				err := client.Request("")
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
			reqResponse, ok := resp.(account.RequestAccountV1Response)
			if ok {
				if &reqResponse.Data != nil {
					for _, a := range reqResponse.Data {
						fmt.Printf("Account id: %d, type: %s, state: %s\n", a.Id, a.Type, a.State)
						if &a.List != nil {
							for _, b := range a.List {
								fmt.Printf("Currency: %s, type: %s, balance: %s\n", b.Currency, b.Type, b.Balance)
							}
						}
					}
				}
			} else {
				fmt.Printf("Received unknown response: %v\n", resp)
			}
		})

	err := client.Connect(false)
	if err != nil {
		fmt.Printf("Client Connect error: %s\n", err)
		return
	}

	fmt.Println("Press ENTER to stop...")
	fmt.Scanln()

	client.Close()
	fmt.Println("Client closed")
}

func subAccountUpdateV1() {
	// Initialize a new instance
	client := new(accountwebsocketclient.SubscribeAccountWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)

	// Set the callback handlers
	client.SetHandler(
		// Connected handler
		func(resp *model.WebSocketV1AuthenticationResponse) {
			if resp.ErrorCode == 0 {
				err := client.Subscribe("1", "1250")
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
			subResponse, ok := resp.(account.SubscribeAccountV1Response)
			if ok {
				if &subResponse.Data != nil {
					fmt.Printf("Account update event: %s\n", subResponse.Data.Event)
					if &subResponse.Data.List != nil {
						for _, b := range subResponse.Data.List {
							fmt.Printf("Account id: %d, currency: %s, type: %s, balance: %s", b.AccountId, b.Currency, b.Type, b.Balance)
						}
					}
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

func subAccountUpdateV2() {
	client := new(accountwebsocketclient.SubscribeAccountWebSocketV2Client).Init(config.AccessKey, config.SecretKey, config.Host)
	client.SetHandler(
		func(resp *model.WebSocketV2AuthenticationResponse) {
			if resp.IsAuth() {
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
		func(resp interface{}) {
			subResponse, ok := resp.(account.SubscribeAccountV2Response)
			if ok {
				if &subResponse.Data != nil {
					b := subResponse.Data
					fmt.Printf("Account id: %d, currency: %s, balance: %s", b.AccountId, b.Currency, b.Balance)
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

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	err = client.UnSubscribe("1", "1250")
	if err != nil {
		fmt.Printf("UnSubscribe error: %s\n", err)
	}

	client.Close()
	fmt.Println("Client closed")
}
