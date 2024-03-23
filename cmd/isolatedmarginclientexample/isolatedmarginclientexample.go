package isolatedmarginclientexample

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/margin"
)

func RunAllExamples() {
	transferIn()
	transferOut()
	getMarginLoanInfo()
	marginOrders()
	marginOrdersRepay()
	marginLoanOrders()
	marginAccountsBalance()
	getMarginLimit()
}

// Transfer specific asset from spot trading account to isolated margin account.
func transferIn() {
	request := margin.IsolatedMarginTransferRequest{
		Currency: "usdt",
		Amount:   "1.0",
		Symbol:   "btcusdt"}
	client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.TransferIn(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		applogger.Info("Data: %+v", resp)
	}
}

// Transfer specific asset from isolated margin account to spot trading account.
func transferOut() {
	request := margin.IsolatedMarginTransferRequest{
		Currency: "usdt",
		Amount:   "1.0",
		Symbol:   "btcusdt"}
	client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.TransferOut(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		applogger.Info("Data: %+v", resp)
	}
}

// Get the loan interest rates and quota applied on the user.
func getMarginLoanInfo() {
	optionalRequest := margin.GetMarginLoanInfoOptionalRequest{Symbols: "btcusdt"}
	client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetMarginLoanInfo(optionalRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, info := range resp {
			applogger.Info("Symbol: %s", info.Symbol)
		}
	}
}

// Place an order to apply a margin loan.
func marginOrders() {
	client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := margin.IsolatedMarginOrdersRequest{
		Currency: "eos",
		Amount:   "0.001",
		Symbol:   "eosht",
	}
	resp, err := client.Apply(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		applogger.Info("Data: %+v", resp)
	}
}

// Repay margin loan with you asset in your margin account.
func marginOrdersRepay() {
	client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	orderId := "12345"
	request := margin.MarginOrdersRepayRequest{Amount: "1.0"}
	resp, err := client.Repay(orderId, request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		applogger.Info("Data: %+v", resp)
	}
}

// Get the margin orders based on a specific searching criteria.
func marginLoanOrders() {
	client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	optionalRequest := margin.IsolatedMarginLoanOrdersOptionalRequest{
		StartDate: "2020-1-1",
	}
	resp, err := client.MarginLoanOrders("btcusdt", optionalRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, order := range resp {
			applogger.Info("Order: %+v", order)
		}
	}
}

// Get the balance of the margin loan account.
func marginAccountsBalance() {
	optionalRequest := margin.MarginAccountsBalanceOptionalRequest{
		Symbol: "btcusdt"}
	client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.MarginAccountsBalance(optionalRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, account := range resp {
			applogger.Info("Id: %d", account.Id)
			for _, balance := range account.List {
				applogger.Info("Balance: %+v", balance)
			}
		}
	}
}

func getMarginLimit() {
	client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetMarginLimit("spot")
	if err != nil {
		applogger.Error("getMarginLimit error: %s", err)
	} else {
		applogger.Info("getMarginLimit, %v", resp)
	}
}
