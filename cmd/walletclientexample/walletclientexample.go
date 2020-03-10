package walletclientexample

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/getrequest"
	"github.com/huobirdcenter/huobi_golang/pkg/postrequest"
)

func RunAllExamples() {
	getDepositAddress()
	getWithdrawQuota()
	createWithdraw()
	cancelWithdraw()
	queryDepositWithdraw()
}

//  Get deposit address of corresponding chain, for a specific crypto currency (except IOTA)
func getDepositAddress() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "usdt"
	resp, err := client.GetDepositAddress(currency)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, result := range resp {
			fmt.Printf("DepositAddress: %+v\n", result)
		}
	}
}

//  Get the withdraw quota for currencies
func getWithdrawQuota() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "usdt"
	resp, err := client.GetWithdrawQuota(currency)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Currency: ", resp.Currency, "Chain: ", resp.Chains[0].Chain, "MaxWithdrawAmt: ", resp.Chains[0].MaxWithdrawAmt)
	}
}

//  Create a withdraw request from your spot trading account to an external address.
func createWithdraw() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	createWithdrawRequest := postrequest.CreateWithdrawRequest{
		Address:  "xxxx",
		Amount:   "1.0",
		Currency: "usdt",
		Fee:      "1.0"}
	resp, err := client.CreateWithdraw(createWithdrawRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

//  Cancel a previously created withdraw request by its transfer id.
func cancelWithdraw() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.CancelWithdraw(12345)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

//  Get all existed withdraws and deposits and return their latest status.
func queryDepositWithdraw() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	depositType := "deposit"
	queryDepositWithdrawOptionalRequest := getrequest.QueryDepositWithdrawOptionalRequest{Currency: "usdt"}
	resp, err := client.QueryDepositWithdraw(depositType, queryDepositWithdrawOptionalRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, result := range resp {
			fmt.Printf("resp: %+v\n", result)
		}
		fmt.Printf("There are total %d deposit-withdraw history", len(resp))
	}
}
