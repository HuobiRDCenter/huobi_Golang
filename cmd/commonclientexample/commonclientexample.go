package commonclientexample

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/getrequest"
)

func RunAllExamples() {
	getSymbols()
	getCurrencys()
	getV2ReferenceCurrencies()
	getTimestamp()
}

func getSymbols() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetSymbols()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, result := range resp {
			fmt.Printf("symbol:%s, BaseCurrency:%s, QuoteCurrency:%s \n", result.Symbol, result.BaseCurrency, result.QuoteCurrency)
		}
		fmt.Printf("There are total %d symbols\n", len(resp))
	}
}

func getCurrencys() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetCurrencys()

	if err != nil {
		fmt.Println(err)
	} else {
		for _, result := range resp {
			fmt.Println(result)
		}
		fmt.Printf("There are total %d currencies\n", len(resp))
	}
}

func getV2ReferenceCurrencies() {
	optionalRequest := getrequest.GetV2ReferenceCurrencies{Currency: "", AuthorizedUser: "true"}

	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetV2ReferenceCurrencies(optionalRequest)

	if err != nil {
		fmt.Println(err)
	} else {
		for _, result := range resp {
			fmt.Printf("currency:%s, ", result.Currency)

			for _, chain := range result.Chains {
				fmt.Printf("Chain:%s, ", chain.Chain)
				fmt.Printf("BaseChain:%s, ", chain.BaseChain)
				fmt.Printf("WithdrawPrecision:%d ", chain.WithdrawPrecision)
			}
			fmt.Printf("\n")
		}

		fmt.Printf("There are total %d CurrencyChain\n", len(resp))
	}
}

func getTimestamp() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetTimestamp()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("timestamp:", resp)
	}
}
