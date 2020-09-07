package subuserclientexample

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/subuser"
	"github.com/shopspring/decimal"
)

func RunAllExamples() {
	createSubUser()
	lockSubUser()
	unlockSubUser()
	setSubUserTradbleMarket()
	setSubUserTransferability()
	subUserTransfer()
	getSubUserDepositAddress()
	querySubUserDepositHistory()
	getSubUserAggregateBalance()
	getSubUserAccount()
}


func createSubUser() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := subuser.CreateSubUserRequest{
		UserList: []subuser.Users{
			subuser.Users{"subuser1412", "sub-user-1-note"},
			subuser.Users{"subuser1413", "sub-user-2-note"},
		},
	}

	resp, err := client.CreateSubUser(request)
	if err != nil {
		applogger.Error("Create sub user error: %s", err)
	} else {
		applogger.Info("Create sub user, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("sub user: %+v", result)
		}
	}
}

func lockSubUser() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	subUserManagementRequest := subuser.SubUserManagementRequest{SubUid: config.SubUid, Action: "lock"}
	resp, err := client.SubUserManagement(subUserManagementRequest)
	if err != nil {
		applogger.Error("Lock sub user error: %s", err)
	} else {
		applogger.Info("Lock sub user: %+v", resp)
	}
}

func unlockSubUser() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	subUserManagementRequest := subuser.SubUserManagementRequest{SubUid: config.SubUid, Action: "unlock"}
	resp, err := client.SubUserManagement(subUserManagementRequest)
	if err != nil {
		applogger.Error("Unlock sub user error: %s", err)
	} else {
		applogger.Info("Unlock sub user: %+v", resp)
	}
}

func setSubUserTradbleMarket() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := subuser.SetSubUserTradableMarketRequest{
		SubUids: config.SubUid,
		AccountType: "isolated-margin",
		Activation: "deactivated",
	}
	resp, err := client.SetSubUserTradableMarket(request)
	if err != nil {
		applogger.Error("Deactivate sub user error: %s", err)
	} else {
		applogger.Info("Deactivate sub user success: %+v", resp)
	}

	request = subuser.SetSubUserTradableMarketRequest{
		SubUids: config.SubUid,
		AccountType: "isolated-margin",
		Activation: "activated",
	}
	resp, err = client.SetSubUserTradableMarket(request)
	if err != nil {
		applogger.Error("Activate sub user error: %s", err)
	} else {
		applogger.Info("Activate sub user: %+v", resp)
	}
}


func setSubUserTransferability() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := subuser.SetSubUserTransferabilityRequest{
		SubUids: config.SubUid,
		AccountType: "spot",
		Transferrable: false,
	}
	resp, err := client.SetSubUserTransferability(request)
	if err != nil {
		applogger.Error("Deactivate sub user error: %s", err)
	} else {
		applogger.Info("Deactivate sub user success: %+v", resp)
	}

	request = subuser.SetSubUserTransferabilityRequest{
		SubUids: config.SubUid,
		AccountType: "spot",
		Transferrable: true,
	}
	resp, err = client.SetSubUserTransferability(request)
	if err != nil {
		applogger.Error("Activate sub user error: %s", err)
	} else {
		applogger.Info("Activate sub user: %+v", resp)
	}
}

func subUserTransfer() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "usdt"
	subUserTransferRequest := subuser.SubUserTransferRequest{
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

func getSubUserDepositAddress() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "usdt"
	resp, err := client.GetSubUserDepositAddress(config.SubUid, currency)
	if err != nil {
		applogger.Error("Get sub user deposit address error: %s", err)
	} else {
		applogger.Info("Get sub user deposit address, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("DepositAddress: %+v", result)
		}
	}
}

func querySubUserDepositHistory() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	optionalRequest := subuser.QuerySubUserDepositHistoryOptionalRequest{Currency: "usdt"}
	resp, err := client.QuerySubUserDepositHistory(config.SubUid, optionalRequest)
	if err != nil {
		applogger.Error("Query sub user deposit history error: %s", err)
	} else {
		applogger.Info("Query sub user deposit history, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("resp: %+v", result)
		}
	}
}

func getSubUserAggregateBalance() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
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

func getSubUserAccount() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
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