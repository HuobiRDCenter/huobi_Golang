package crossmarginclientexample

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/getrequest"
	"github.com/huobirdcenter/huobi_golang/pkg/postrequest"
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

//  Transfer specific asset from spot trading account to cross margin account.
func transferIn() {
	request := postrequest.CrossMarginTransferRequest{
		Currency: "usdt",
		Amount:   "1.0"}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.TransferIn(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data: ", resp)
	}
}

//  Transfer specific asset from cross margin account to spot trading account.
func transferOut() {
	request := postrequest.CrossMarginTransferRequest{
		Currency: "usdt",
		Amount:   "1.0"}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.TransferOut(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data: ", resp)
	}
}

//  Get the loan interest rates and loan quota applied on the user.
func getMarginLoanInfo() {
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetMarginLoanInfo()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, info := range resp {
			fmt.Printf("Info: %+v \n", info)
		}
	}
}

//  Place an order to apply a margin loan.
func marginOrders() {
	request := postrequest.CrossMarginOrdersRequest{Currency: "usdt",
		Amount: "1.0"}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.ApplyLoan(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data: ", resp)
	}
}

//  Repays margin loan with you asset in your margin account.
func marginOrdersRepay() {
	orderId := "12345"
	request := postrequest.MarginOrdersRepayRequest{Amount: "1.0"}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	err := client.Repay(orderId, request)
	if err != nil {
		fmt.Println(err)
	}
}

//  Get the margin orders based on a specific searching criteria.
func marginLoanOrders() {
	optionalRequest := getrequest.CrossMarginLoanOrdersOptionalRequest{}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.MarginLoanOrders(optionalRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, order := range resp {
			fmt.Printf("Order: %+v\n", order)
		}
	}
}

//  Get the balance of the margin loan account.
func marginAccountsBalance() {
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.MarginAccountsBalance("")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, account := range resp.List {
			fmt.Printf("Account: %+v\n", account)
		}
	}
}
