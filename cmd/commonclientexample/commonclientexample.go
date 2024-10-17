package commonclientexample

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/common"
)

func RunAllExamples() {
	getSystemStatus()
	getMarketStatus()
	getSymbols()
	getCurrencys()
	getV2ReferenceCurrencies()
	getTimestamp()
	getSymbolsV2()
	getCurrenciesV2()
	getCurrencysV1()
	getSymbolsV1()
	getMarketSymbols()
	getChains()
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

func getMarketStatus() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetMarketStatus()
	if err != nil {
		applogger.Error("Get market status error: %s", err)
	} else {
		applogger.Info("Get market status, status: %d", resp.MarketStatus)
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
			applogger.Info("symbol: %s, %+v", result.Symbol, result)
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
	optionalRequest := common.GetV2ReferenceCurrencies{Currency: "", AuthorizedUser: "true"}

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

func getSymbolsV2() {
	client := new(client.CommonClient).Init(config.Host)
	optionalRequest := common.GetSymbolsV2{}
	resp, err := client.GetSymbolsV2(optionalRequest)
	if err != nil {
		applogger.Error("getSymbolsV2 error: %s", err)
	} else {
		applogger.Info("getSymbolsV2, %v", resp)
	}
}

func getCurrenciesV2() {
	client := new(client.CommonClient).Init(config.Host)
	optionalRequest := common.GetCurrenciesV2{}
	resp, err := client.GetCurrenciesV2(optionalRequest)
	if err != nil {
		applogger.Error("getCurrenciesV2 error: %s", err)
	} else {
		applogger.Info("getCurrenciesV2, %v", resp)
	}
}

func getCurrencysV1() {
	client := new(client.CommonClient).Init(config.Host)
	optionalRequest := common.GetCurrencysV1{}
	resp, err := client.GetCurrencysV1(optionalRequest)
	if err != nil {
		applogger.Error("getSymbolsV2 error: %s", err)
	} else {
		applogger.Info("getSymbolsV2, %v", resp)
	}
}

func getSymbolsV1() {
	client := new(client.CommonClient).Init(config.Host)
	optionalRequest := common.GetSymbolsV1{}
	resp, err := client.GetSymbolsV1(optionalRequest)
	if err != nil {
		applogger.Error("GetSymbolsV1 error: %s", err)
	} else {
		applogger.Info("GetSymbolsV1, %v", resp)
	}
}

func getMarketSymbols() {
	client := new(client.CommonClient).Init(config.Host)
	optionalRequest := common.GetMarketSymbols{}
	resp, err := client.GetMarketSymbols(optionalRequest)
	if err != nil {
		applogger.Error("GetMarketSymbols error: %s", err)
	} else {
		applogger.Info("GetMarketSymbols, %v", resp)
	}
}

func getChains() {
	client := new(client.CommonClient).Init(config.Host)
	optionalRequest := common.GetChains{}
	resp, err := client.GetChains(optionalRequest)
	if err != nil {
		applogger.Error("GetChains error: %s", err)
	} else {
		applogger.Info("GetChains, %v", resp)
	}
}
