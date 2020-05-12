package commonclientexample

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/getrequest"
)

func RunAllExamples() {
	getSystemStatus()
	getSymbols()
	getCurrencys()
	getV2ReferenceCurrencies()
	getTimestamp()
}

func getSystemStatus() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetSystemStatus()
	if err != nil {
		applogger.Error("Get system status error: %s", err)
	} else {
		applogger.Info("Get system status %s", resp)
	}
}

func getSymbols() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetSymbols()
	if err != nil {
		applogger.Error("Get symbols error: %s", err)
	} else {
		applogger.Info("Get symbols, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("symbol=%s, BaseCurrency=%s, QuoteCurrency=%s", result.Symbol, result.BaseCurrency, result.QuoteCurrency)
		}
	}
}

func getCurrencys() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetCurrencys()

	if err != nil {
		applogger.Error("Get currency error: %s", err)
	} else {
		applogger.Info("Get currency, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("currency: %+v", result)
		}
	}
}

func getV2ReferenceCurrencies() {
	optionalRequest := getrequest.GetV2ReferenceCurrencies{Currency: "", AuthorizedUser: "true"}

	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetV2ReferenceCurrencies(optionalRequest)

	if err != nil {
		applogger.Error("Get reference currency error: %s", err)
	} else {
		applogger.Info("Get reference currency, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("currency:%s, ", result.Currency)

			for _, chain := range result.Chains {
				applogger.Info("Chain: %+v", chain)
			}
		}
	}
}

func getTimestamp() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetTimestamp()

	if err != nil {
		applogger.Error("Get timestamp error: %s", err)
	} else {
		applogger.Info("Get timestamp: %d", resp)
	}
}
