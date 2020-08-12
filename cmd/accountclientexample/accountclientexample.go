package accountclientexample

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/account"
	"github.com/huobirdcenter/huobi_golang/pkg/postrequest"
	"github.com/shopspring/decimal"
)

func RunAllExamples() {
	getAccountInfo()
	getAccountBalance()
	getAccountHistory()
	getAccountLedger()
	transferFromFutureToSpot()
	transferFromSpotToFuture()
	subUserTransfer()
	getSubUserAggregateBalance()
	getSubUserAccount()
	lockSubUser()
	unlockSubUser()
}

func getAccountInfo() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetAccountInfo()
	if err != nil {
		applogger.Error("Get account error: %s", err)
	} else {
		applogger.Info("Get account, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("account: %+v", result)
		}

	}
}

func getAccountBalance() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetAccountBalance(config.AccountId)
	if err != nil {
		applogger.Error("Get account balance error: %s", err)
	} else {
		applogger.Info("Get account balance: id=%d, type=%s, state=%s, count=%d",
			resp.Id, resp.Type, resp.State, len(resp.List))
		if resp.List != nil {
			for _, result := range resp.List {
				applogger.Info("Account balance: %+v", result)
			}
		}
	}
}

func getAccountHistory() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	getAccountHistoryOptionalRequest := account.GetAccountHistoryOptionalRequest{}
	resp, err := client.GetAccountHistory(config.AccountId, getAccountHistoryOptionalRequest)
	if err != nil {
		applogger.Error("Get account history error: %s", err)
	} else {
		applogger.Info("Get account history, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("Account history: %+v", result)
		}
	}
}

func getAccountLedger() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	getAccountLedgerOptionalRequest := account.GetAccountLedgerOptionalRequest{}
	resp, err := client.GetAccountLedger(config.AccountId, getAccountLedgerOptionalRequest)
	if err != nil {
		applogger.Error("Get account ledger error: %s", err)
	} else {
		applogger.Info("Get account ledger, count=%d", len(resp))
		for _, l := range resp {
			applogger.Info("Account legder: AccountId: %d, Currency: %s, Amount: %v, Transferer: %d, Transferee: %d", l.AccountId, l.Currency, l.TransactAmt, l.Transferer, l.Transferee)
		}
	}
}

func getSubUserAccount() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetSubUserAccount(config.SubUid)
	if err != nil {
		applogger.Error("Get sub user account error: %s", err)
	} else {
		applogger.Info("Get sub user account, count=%d", len(resp))
		for _, account := range resp {
			applogger.Info("account id: %d, type: %s, currency count=%d", account.Id, account.Type, len(account.List))
			if account.List != nil {
				for _, currency := range account.List {
					applogger.Info("currency: %+v", currency)
				}
			}
		}
	}
}

func transferFromFutureToSpot() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	futuresTransferRequest := postrequest.FuturesTransferRequest{Currency: "btc", Amount: decimal.NewFromFloat(0.001), Type: "futures-to-pro"}
	resp, err := client.FuturesTransfer(futuresTransferRequest)
	if err != nil {
		applogger.Error("Transfer from future to spot error: %s", err)
	} else {
		applogger.Info("Transfer from future to spot success: id=%d", resp)
	}
}

func transferFromSpotToFuture() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	futuresTransferRequest := postrequest.FuturesTransferRequest{Currency: "btc", Amount: decimal.NewFromFloat(0.001), Type: "pro-to-futures"}
	resp, err := client.FuturesTransfer(futuresTransferRequest)
	if err != nil {
		applogger.Error("Transfer from spot to future error: %s", err)
	} else {
		applogger.Info("Transfer from spot to future success: id=%d", resp)
	}
}

func getSubUserAggregateBalance() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetSubUserAggregateBalance()
	if err != nil {
		applogger.Error("Get sub user aggregated balance error: %s", err)
	} else {
		applogger.Info("Get sub user aggregated balance, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("balance: %+v", result)
		}
	}
}

func subUserTransfer() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "usdt"
	subUserTransferRequest := postrequest.SubUserTransferRequest{
		SubUid:   config.SubUid,
		Currency: currency,
		Amount:   decimal.NewFromInt(1),
		Type:     "master-transfer-in",
	}
	resp, err := client.SubUserTransfer(subUserTransferRequest)
	if err != nil {
		applogger.Error("Transfer error: %s", err)
	} else {
		applogger.Info("Transfer successfully, id=%s", resp)

	}
}

func lockSubUser() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	subUserManagementRequest := postrequest.SubUserManagementRequest{SubUid: config.SubUid, Action: "lock"}
	resp, err := client.SubUserManagement(subUserManagementRequest)
	if err != nil {
		applogger.Error("Lock sub user error: %s", err)
	} else {
		applogger.Info("Lock sub user: %+v", resp)
	}
}

func unlockSubUser() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	subUserManagementRequest := postrequest.SubUserManagementRequest{SubUid: config.SubUid, Action: "unlock"}
	resp, err := client.SubUserManagement(subUserManagementRequest)
	if err != nil {
		applogger.Error("Unlock sub user error: %s", err)
	} else {
		applogger.Info("Unlock sub user: %+v", resp)
	}
}
