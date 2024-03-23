package accountclientexample

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/account"
	"github.com/shopspring/decimal"
)

func RunAllExamples() {
	getAccountInfo()
	getAccountBalance()
	getAccountAssetValuation()
	transferAccount()
	getAccountHistory()
	getAccountLedger()
	transferFromFutureToSpot()
	transferFromSpotToFuture()
	getPointBalance()
	transferPoint()
	getValuation()
	transfer()
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

func getAccountAssetValuation() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetAccountAssetValuation("spot", "USD", 0)
	if err != nil {
		applogger.Error("Get account asset valuation error: %s", err)
	} else {
		applogger.Info("Get account asset valuation, balance: %s, timestamp: %d", resp.Data.Balance, resp.Data.Timestamp)
	}
}

func transferAccount() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := account.TransferAccountRequest{
		FromUser:        125753978,
		FromAccountType: "spot",
		FromAccount:     11136102,
		ToUser:          128654510,
		ToAccountType:   "spot",
		ToAccount:       12825690,
		Currency:        "ht",
		Amount:          "0.18",
	}
	resp, err := client.TransferAccount(request)
	if err != nil {
		applogger.Error("Transfer account error: %s", err)
	} else {
		applogger.Info("Transfer account, %v", resp.Data)
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

func transferFromFutureToSpot() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	futuresTransferRequest := account.FuturesTransferRequest{Currency: "btc", Amount: decimal.NewFromFloat(0.001), Type: "futures-to-pro"}
	resp, err := client.FuturesTransfer(futuresTransferRequest)
	if err != nil {
		applogger.Error("Transfer from future to spot error: %s", err)
	} else {
		applogger.Info("Transfer from future to spot success: id=%d", resp)
	}
}

func transferFromSpotToFuture() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	futuresTransferRequest := account.FuturesTransferRequest{Currency: "btc", Amount: decimal.NewFromFloat(0.001), Type: "pro-to-futures"}
	resp, err := client.FuturesTransfer(futuresTransferRequest)
	if err != nil {
		applogger.Error("Transfer from spot to future error: %s", err)
	} else {
		applogger.Info("Transfer from spot to future success: id=%d", resp)
	}
}

func getPointBalance() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetPointBalance(config.SubUids)
	if err != nil {
		applogger.Error("Get point balance error: %s", err)
	} else {
		applogger.Info("Get point balance: id=%s, balance=%s, state=%s, count=%d",
			resp.Data.AccountId, resp.Data.AccountBalance, resp.Data.AccountStatus, len(resp.Data.GroupIds))
		if resp.Data.GroupIds != nil {
			for _, result := range resp.Data.GroupIds {
				applogger.Info("Account balance: %+v", result)
			}
		}
	}
}

func transferPoint() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := account.TransferPointRequest{FromUid: "125753978", ToUid: "128654685", GroupId: 0, Amount: "0"}
	resp, err := client.TransferPoint(request)
	if err != nil {
		applogger.Error("Transfer points error: %s", err)
	} else {
		applogger.Info("Transfer point success: id=%s, time=%d", resp.Data.TransactId, resp.Data.TransactTime)
	}
}

func getValuation() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	optionalRequest := account.GetValuation{ValuationCurrency: "BTC"}
	resp, err := client.GetValuation("spot", optionalRequest)
	if err != nil {
		applogger.Error("Get Valuation error: %s", err)
	} else {
		applogger.Info("Valuation, %v", resp.Data)
	}
}

func transfer() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := account.TransferRequest{From: "spot", To: "linear-swap", Currency: "usdt", Amount: 100, MarginAccount: "USDT"}
	resp, err := client.Transfer(request)
	if err != nil {
		applogger.Error("transfer error: %s", err)
	} else {
		applogger.Info("transfer, %v", resp.Data)
	}
}
