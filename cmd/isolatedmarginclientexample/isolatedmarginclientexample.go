package isolatedmarginclientexample

import (
	"../../config"
	"../../pkg/client"
	"../../pkg/getrequest"
	"../../pkg/postrequest"
	"fmt"
)

func RunAllExamples() {
	transferIn()
	transferOut()
	getMarginLoanInfo()
	marginOrders()
	marginOrdersRepay()
	marginLoanOrders()
	marginAccountsBalance()
}

// Transfer specific asset from spot trading account to isolated margin account.
func transferIn() {
	request := postrequest.IsolatedMarginTransferRequest{
		Currency: "usdt",
		Amount:   "1.0",
		Symbol:   "btcusdt"}
	client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.TransferIn(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data: ", resp)
	}
}

//  Transfer specific asset from isolated margin account to spot trading account.
func transferOut() {
	request := postrequest.IsolatedMarginTransferRequest{
		Currency: "usdt",
		Amount:   "1.0",
		Symbol:   "btcusdt"}
	client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.TransferOut(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data: ", resp)
	}
}

//  Get the loan interest rates and quota applied on the user.
func getMarginLoanInfo() {
	optionalRequest := getrequest.GetMarginLoanInfoOptionalRequest{Symbols: "btcusdt"}
	client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetMarginLoanInfo(optionalRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, info := range resp {
			fmt.Println("Symbol: ", info.Symbol)

		}
	}
}

//  Place an order to apply a margin loan.
func marginOrders() {
	client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := postrequest.IsolatedMarginOrdersRequest{
		Currency: "eos",
		Amount: "0.001",
		Symbol: "eosht",
	}
	resp, err := client.MarginOrders(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data: ", resp)
	}
}

// Repay margin loan with you asset in your margin account.
func marginOrdersRepay() {
	client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	orderId := "12345"
	request := postrequest.MarginOrdersRepayRequest{Amount: "1.0"}
	resp, err := client.MarginOrdersRepay(orderId, request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data: ", resp)
	}
}

//  Get the margin orders based on a specific searching criteria.
func marginLoanOrders() {
	client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	optionalRequest := getrequest.IsolatedMarginLoanOrdersOptionalRequest{
		StartDate: "2020-1-1",
	}
	resp, err := client.MarginLoanOrders("btcusdt", optionalRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, order := range resp {
			fmt.Println("Order: ", order)
		}
	}
}

//  Get the balance of the margin loan account.
func marginAccountsBalance() {
	optionalRequest := getrequest.MarginAccountsBalanceOptionalRequest{
		Symbol: "btcusdt"}
	client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.MarginAccountsBalance(optionalRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, account := range resp {
			fmt.Println("Id: ", account.Id)
			for _, balance := range account.List {
				fmt.Printf("Balance: %+v\n", balance)
			}
		}
	}
}
