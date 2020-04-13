package accountclientexample

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/getrequest"
	"github.com/huobirdcenter/huobi_golang/pkg/postrequest"
	"github.com/shopspring/decimal"
)

func RunAllExamples() {
	getAccountInfo()
	getAccountBalance()
	getAccountHistory()
	getAccountLedger()
	getSubUserAccount()
	subUserManagement()
	futuresTransfer()
	getSubUserAggregateBalance()
	subUserTransfer()
}

//  Get a list of accounts owned by this API user.
func getAccountInfo() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetAccountInfo()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, result := range resp {
			fmt.Printf("account: %+v\n", result)
		}

	}
}

//   Get the balance of an account specified by account id.
func getAccountBalance() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetAccountBalance(config.AccountId)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Id: ", resp.Id, "Type: ", resp.Type, "State: ", resp.State)
	}
}

//  Get the amount changes of specified user's account.
func getAccountHistory() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	getAccountHistoryOptionalRequest := getrequest.GetAccountHistoryOptionalRequest{}
	resp, err := client.GetAccountHistory(config.AccountId, getAccountHistoryOptionalRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, result := range resp {
			fmt.Println(result)
		}
	}
}

func getAccountLedger() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	getAccountLedgerOptionalRequest := getrequest.GetAccountLedgerOptionalRequest{}
	resp, err := client.GetAccountLedger(config.AccountId, getAccountLedgerOptionalRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, l := range resp {
			fmt.Printf("Account legder: AccountId: %d, Currency: %s, Amount: %v, Transferer: %d, Transferee: %d\n", l.AccountId, l.Currency, l.TransactAmt, l.Transferer, l.Transferee)
		}
	}
}

//  Get the balance of a sub-account specified by sub-uid.
func getSubUserAccount() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetSubUserAccount(config.SubUid)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, result := range resp {
			fmt.Println("Id: ", result.Id, "Type: ", result.Type)
		}

	}
}

//  This func allows parent user to lock or unlock a specific sub user.
func subUserManagement() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	subUserManagementRequest := postrequest.SubUserManagementRequest{SubUid: config.SubUid, Action: "unlock"}
	resp, err := client.SubUserManagement(subUserManagementRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("resp: %+v\n", resp)
	}
}

//  Transfer fund between spot account and future contract account.
func futuresTransfer() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	futuresTransferRequest := postrequest.FuturesTransferRequest{Currency: "btc", Amount: decimal.NewFromFloat(0.001), Type: "futures-to-pro"}
	resp, err := client.FuturesTransfer(futuresTransferRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

// Get the balances of all the sub-account aggregated.
func getSubUserAggregateBalance() {
	client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetSubUserAggregateBalance()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, result := range resp {
			fmt.Printf("result: %+v\n", result)
		}
	}
}

// Transfer asset between parent and sub account.
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
		fmt.Println(err)
	} else {
		fmt.Println(resp)

	}
}
