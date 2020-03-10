package etfclientexample

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/postrequest"
)

func RunAllExamples() {
	getSwapConfig()
	getSwapList()
	swapIn()
	swapOut()
}

//  Get the basic information of ETF creation and redemption
func getSwapConfig() {
	client := new(client.ETFClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "hb10"
	resp, err := client.GetSwapConfig(currency)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("EtfStatus: ", resp.EtfStatus, "PurchaseFeeRate: ", resp.PurchaseFeeRate)
	}
}

//  Get past creation and redemption.(up to 100 records)
func getSwapList() {
	client := new(client.ETFClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "hb10"
	resp, err := client.GetSwapList(currency, 0, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, result := range resp {
			fmt.Printf("SwapList: %+v\n", result)
		}
	}
}

//  Allow clients to order creation of ETF.
func swapIn() {
	client := new(client.ETFClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "hb10"
	swapRequest := postrequest.SwapRequest{
		EtfName: currency,
		Amount:  10,
	}
	resp, err := client.SwapIn(swapRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("isSucceed: ", resp)
	}
}

//  Allow clients to order redemption of ETF.
func swapOut() {
	client := new(client.ETFClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "hb10"
	swapRequest := postrequest.SwapRequest{
		EtfName: currency,
		Amount:  10,
	}
	resp, err := client.SwapOut(swapRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("isSucceed: ", resp)
	}
}
