package walletclientexample

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/wallet"
)

func RunAllExamples() {
	getDepositAddress()
	getWithdrawQuota()
	getWithdrawAddress()
	createWithdraw()
	cancelWithdraw()
	queryDepositWithdraw()
	getWithdraw()
}

func getDepositAddress() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	currency := "usdt"
	resp, err := client.GetDepositAddress(currency)
	if err != nil {
		applogger.Error("Get deposit address error: %s", err)
	} else {
		applogger.Info("Get deposit address, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("DepositAddress: %+v", result)
		}
	}
}

func getWithdrawQuota() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	currency := "usdt"
	resp, err := client.GetWithdrawQuota(currency)
	if err != nil {
		applogger.Error("Get withdraw quota error: %s", err)
	} else {
		applogger.Info("Currency: %s, Chain: %s, MaxWithdrawAmt: %s", resp.Currency, resp.Chains[0].Chain, resp.Chains[0].MaxWithdrawAmt)
	}
}

func getWithdrawAddress() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	request := new(model.GetRequest).Init()
	request.AddParam("currency", "btc")

	resp, err := client.GetWithdrawAddress(request)
	if err != nil {
		applogger.Error("Get withdraw address error: %s", err)
	} else {
		if resp.Data != nil {
			applogger.Info("Get withdraw address, count=%d", len(resp.Data))
			for _, addr := range resp.Data {
				applogger.Info("Get withdraw address, Currency: %s, Chain: %s, Address: %s, AddressTag: %s", addr.Currency, addr.Chain, addr.Address, addr.AddressTag)
			}
		}
	}
}

func createWithdraw() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	createWithdrawRequest := wallet.CreateWithdrawRequest{
		Address:  "xxxx",
		Amount:   "1.0",
		Currency: "usdt",
		Fee:      "1.0"}
	resp, err := client.CreateWithdraw(createWithdrawRequest)
	if err != nil {
		applogger.Error("Create withdraw request error: %s", err)
	} else {
		applogger.Info("Create withdraw request successfully: id=%d", resp)
	}
}

func cancelWithdraw() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp, err := client.CancelWithdraw(12345)
	if err != nil {
		applogger.Error("Cancel withdraw error: %s", err)
	} else {
		applogger.Info("Cancel withdraw successfully: id=%d", resp)
	}
}

func queryDepositWithdraw() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	depositType := "deposit"
	queryDepositWithdrawOptionalRequest := wallet.QueryDepositWithdrawOptionalRequest{Currency: "usdt"}
	resp, err := client.QueryDepositWithdraw(depositType, queryDepositWithdrawOptionalRequest)
	if err != nil {
		applogger.Error("Query deposit-withdraw history error: %s", err)
	} else {
		applogger.Info("Query deposit-withdraw history, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("resp: %+v", result)
		}
	}
}

func getWithdraw() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp, err := client.GetWithdraw("1113")
	if err != nil {
		applogger.Error("getWithdraw error: %s", err)
	} else {
		applogger.Info("getWithdraw, %v", resp.Data)
	}
}
