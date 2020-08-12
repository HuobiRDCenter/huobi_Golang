package subuserclientexample

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/subuser"
)

func RunAllExamples() {
	createSubUser()
	getSubUserDepositAddress()
	querySubUserDepositHistory()
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
