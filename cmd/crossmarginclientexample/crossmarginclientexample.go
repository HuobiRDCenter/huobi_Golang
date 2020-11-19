package crossmarginclientexample

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
	genernalMarginOrdersRepay()
	genernalMarginLoanOrders()
}

//  Transfer specific asset from spot trading account to cross margin account.
func transferIn() {
	request := margin.CrossMarginTransferRequest{
		Currency: "usdt",
		Amount:   "1.0"}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.TransferIn(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		applogger.Info("Data: %+v", resp)
	}
}

//  Transfer specific asset from cross margin account to spot trading account.
func transferOut() {
	request := margin.CrossMarginTransferRequest{
		Currency: "usdt",
		Amount:   "1.0"}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.TransferOut(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		applogger.Info("Data: %+v", resp)
	}
}

//  Get the loan interest rates and loan quota applied on the user.
func getMarginLoanInfo() {
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetMarginLoanInfo()
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, info := range resp {
			applogger.Info("Info: %+v", info)
		}
	}
}

//  Place an order to apply a margin loan.
func marginOrders() {
	request := margin.CrossMarginOrdersRequest{Currency: "usdt",
		Amount: "1.0"}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.ApplyLoan(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		applogger.Info("Data: %+v", resp)
	}
}

//  Repays margin loan with you asset in your margin account.
func marginOrdersRepay() {
	orderId := "12345"
	request := margin.MarginOrdersRepayRequest{Amount: "1.0"}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.Repay(orderId, request)
	if err != nil {
		applogger.Error("Repay error: %s", err)
	} else {
		applogger.Info("Repay successfully, id=%d", resp)
	}
}

//  Get the margin orders based on a specific searching criteria.
func marginLoanOrders() {
	optionalRequest := margin.CrossMarginLoanOrdersOptionalRequest{}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.MarginLoanOrders(optionalRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, order := range resp {
			applogger.Info("Order: %+v", order)
		}
	}
}

//  Get the balance of the margin loan account.
func marginAccountsBalance() {
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.MarginAccountsBalance("")
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, account := range resp.List {
			applogger.Info("Account: %+v", account)
		}
	}
}

//  Repays general margin loan with you asset in your margin account.
func genernalMarginOrdersRepay() {
	request := margin.CrossMarginGeneralReplayLoanOptionalRequest{AccountId: "12345", Currency:"btc", Amount:"0.01"}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GeneralRepay(request)
	if err != nil {
		applogger.Error("Repay error: %s", err)
	} else {
		for _, order := range resp {
			applogger.Info("Repay Order:%+v", order)
		}
	}
}

//  Get the genernal margin orders based on a specific searching criteria.
func genernalMarginLoanOrders() {
	optionalRequest := margin.CrossMarginGeneralReplayLoanRecordsOptionalRequest{}
	client := new(client.CrossMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GeneralMarginLoanOrders(optionalRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, record := range resp {
			applogger.Info("Record: RepayId:%v, RepayTime:%v, AccountId:%v, Currency:%s, RepaidAmount:%s, TransactId:%+v", record.RepayId, record.RepayTime, record.AccountId, record.Currency, record.RepaidAmount, record.TransactIds)
		}
	}
}
