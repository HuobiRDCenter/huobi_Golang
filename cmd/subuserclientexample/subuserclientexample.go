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
	getUid()
	deductMode()
	getApiKey()
	getUserList()
	getUserState()
	getAccountList()
	apiKeyGeneration()
	apiKeyModification()
	apiKeyDeletion()
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
		SubUids:     config.SubUids,
		AccountType: "isolated-margin",
		Activation:  "deactivated",
	}
	resp, err := client.SetSubUserTradableMarket(request)
	if err != nil {
		applogger.Error("Deactivate sub user error: %s", err)
	} else {
		for _, result := range resp {
			applogger.Info("Deactivate sub user success: %+v", result)
		}
	}

	request = subuser.SetSubUserTradableMarketRequest{
		SubUids:     config.SubUids,
		AccountType: "isolated-margin",
		Activation:  "activated",
	}
	resp, err = client.SetSubUserTradableMarket(request)
	if err != nil {
		applogger.Error("Activate sub user error: %s", err)
	} else {
		for _, result := range resp {
			applogger.Info("Activate sub user success: %+v", result)
		}
	}
}

func setSubUserTransferability() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := subuser.SetSubUserTransferabilityRequest{
		SubUids:       config.SubUids,
		AccountType:   "spot",
		Transferrable: false,
	}
	resp, err := client.SetSubUserTransferability(request)
	if err != nil {
		applogger.Error("Deactivate sub user error: %s", err)
	} else {
		for _, result := range resp {
			applogger.Info("Deactivate sub user success: %+v", result)
		}
	}

	request = subuser.SetSubUserTransferabilityRequest{
		SubUids:       config.SubUids,
		AccountType:   "spot",
		Transferrable: true,
	}
	resp, err = client.SetSubUserTransferability(request)
	if err != nil {
		applogger.Error("Activate sub user error: %s", err)
	} else {
		for _, result := range resp {
			applogger.Info("Activate sub user: %+v", result)
		}
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

func getUid() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetUid()
	if err != nil {
		applogger.Error("Get uid error: %s", err)
	} else {
		applogger.Info("Get uid: %d", resp)
	}
}

func deductMode() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := subuser.DeductModeRequest{SubUids: 178211, DeductMode: "master"}
	resp, err := client.DeductMode(request)
	if err != nil {
		applogger.Error("deductMode error: %s", err)
	} else {
		applogger.Info("deductMode, %v", resp.Data)
	}
}

func getApiKey() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := subuser.GetApiKey{AccessKey: config.AccessKey}
	resp, err := client.GetApiKey(config.SubUid, request)
	if err != nil {
		applogger.Error("GetApiKey error: %s", err)
	} else {
		applogger.Info("GetApiKey, %v", resp)
	}
}

func getUserList() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := subuser.GetUserList{}
	resp, err := client.GetUserList(request)
	if err != nil {
		applogger.Error("GetUserList error: %s", err)
	} else {
		applogger.Info("GetUserList, %v", resp)
	}
}

func getUserState() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetUserState(config.SubUid)
	if err != nil {
		applogger.Error("GetUserState error: %s", err)
	} else {
		applogger.Info("GetUserState, %v", resp)
	}
}

func getAccountList() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetAccountList(config.SubUid)
	if err != nil {
		applogger.Error("GetAccountList error: %s", err)
	} else {
		applogger.Info("GetAccountList, %v", resp)
	}
}

func apiKeyGeneration() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := subuser.ApiKeyGenerationRequest{OtpToken: "178211", SubUid: config.SubUid, Note: "178211", Permission: "178211", IPAddresses: "178211"}
	resp, err := client.ApiKeyGeneration(request)
	if err != nil {
		applogger.Error("ApiKeyGeneration error: %s", err)
	} else {
		applogger.Info("ApiKeyGeneration, %v", resp)
	}
}

func apiKeyModification() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := subuser.ApiKeyModificationRequest{SubUid: 178211, AccessKey: "178211", Note: "178211", Permission: "178211", IPAddresses: "178211"}
	resp, err := client.ApiKeyModification(request)
	if err != nil {
		applogger.Error("ApiKeyModification error: %s", err)
	} else {
		applogger.Info("ApiKeyModification, %v", resp)
	}
}

func apiKeyDeletion() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := subuser.ApiKeyDeletionRequest{SubUid: 178211, AccessKey: "178211"}
	resp, err := client.ApiKeyDeletion(request)
	if err != nil {
		applogger.Error("ApiKeyDeletion error: %s", err)
	} else {
		applogger.Info("ApiKeyDeletion, %v", resp)
	}
}
