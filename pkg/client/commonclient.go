package client

import (
	"encoding/json"
	"errors"
	"github.com/huobirdcenter/huobi_golang/internal"
	"github.com/huobirdcenter/huobi_golang/internal/requestbuilder"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/common"
)

// Responsible to get common information
type CommonClient struct {
	publicUrlBuilder *requestbuilder.PublicUrlBuilder
}

// Initializer
func (p *CommonClient) Init(host string) *CommonClient {
	p.publicUrlBuilder = new(requestbuilder.PublicUrlBuilder).Init(host)
	return p
}

func (p *CommonClient) GetSystemStatus() (string, error) {
	url := "https://status.huobigroup.com/api/v2/summary.json"
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return "", getErr
	}

	return getResp, nil
}

// Returns current market status
func (p *CommonClient) GetMarketStatus() (*common.MarketStatus, error) {
	url := p.publicUrlBuilder.Build("/v2/market-status", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := common.GetMarketStatusResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && &result.Data != nil {
		return &result.Data, nil
	}
	return nil, errors.New(getResp)
}

// Get all Supported Trading Symbol
// This endpoint returns all Huobi's supported trading symbol.
func (p *CommonClient) GetSymbols() ([]common.Symbol, error) {
	url := p.publicUrlBuilder.Build("/v1/common/symbols", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := common.GetSymbolsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}

// Get all Supported Currencies
// This endpoint returns all Huobi's supported trading currencies.
func (p *CommonClient) GetCurrencys() ([]string, error) {
	url := p.publicUrlBuilder.Build("/v1/common/currencys", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := common.GetCurrenciesResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)

	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}

// APIv2 - Currency & Chains
// API user could query static reference information for each currency, as well as its corresponding chain(s). (Public Endpoint)
func (p *CommonClient) GetV2ReferenceCurrencies(optionalRequest common.GetV2ReferenceCurrencies) ([]common.CurrencyChain, error) {
	request := new(model.GetRequest).Init()
	if optionalRequest.Currency != "" {
		request.AddParam("currency", optionalRequest.Currency)
	}
	if optionalRequest.AuthorizedUser != "" {
		request.AddParam("authorizedUser", optionalRequest.AuthorizedUser)
	}

	url := p.publicUrlBuilder.Build("/v2/reference/currencies", request)

	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := common.GetV2ReferenceCurrenciesResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)

	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}

	return nil, errors.New(result.Message)
}

// Get Current Timestamp
// This endpoint returns the current timestamp, i.e. the number of milliseconds that have elapsed since 00:00:00 UTC on 1 January 1970.
func (p *CommonClient) GetTimestamp() (int, error) {
	url := p.publicUrlBuilder.Build("/v1/common/timestamp", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return 0, getErr
	}

	result := common.GetTimestampResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)

	if jsonErr != nil {
		return 0, jsonErr
	}
	if result.Status == "ok" && result.Data != 0 {
		return result.Data, nil
	}
	return 0, errors.New(getResp)
}

// 获取所有交易对(V2)
func (p *CommonClient) GetSymbolsV2(optionalRequest common.GetSymbolsV2) ([]common.SymbolV2, error) {
	request := new(model.GetRequest).Init()
	if optionalRequest.Ts != "" {
		request.AddParam("ts", optionalRequest.Ts)
	}

	url := p.publicUrlBuilder.Build("/v2/settings/common/symbols", request)

	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := common.GetSymbolsV2Response{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)

	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// 获取所有币种(V2)
func (p *CommonClient) GetCurrenciesV2(optionalRequest common.GetCurrenciesV2) ([]common.CurrenciesV2, error) {
	request := new(model.GetRequest).Init()
	if optionalRequest.Ts != "" {
		request.AddParam("ts", optionalRequest.Ts)
	}

	url := p.publicUrlBuilder.Build("/v2/settings/common/currencies", request)

	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := common.GetCurrenciesV2Response{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)

	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// 获取币种配置
func (p *CommonClient) GetCurrencysV1(optionalRequest common.GetCurrencysV1) ([]common.CurrencysV1, error) {
	request := new(model.GetRequest).Init()
	if optionalRequest.Ts != "" {
		request.AddParam("ts", optionalRequest.Ts)
	}

	url := p.publicUrlBuilder.Build("/v1/settings/common/currencys", request)

	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := common.GetCurrencysV1Response{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)

	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// 获取交易对配置
func (p *CommonClient) GetSymbolsV1(optionalRequest common.GetSymbolsV1) ([]common.SymbolsV1, error) {
	request := new(model.GetRequest).Init()
	if optionalRequest.Ts != "" {
		request.AddParam("ts", optionalRequest.Ts)
	}

	url := p.publicUrlBuilder.Build("/v1/settings/common/symbols", request)

	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := common.GetSymbolsV1Response{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)

	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// 获取市场交易对配置
func (p *CommonClient) GetMarketSymbols(optionalRequest common.GetMarketSymbols) ([]common.MarketSymbols, error) {
	request := new(model.GetRequest).Init()
	if optionalRequest.Symbols != "" {
		request.AddParam("symbols", optionalRequest.Symbols)
	}
	if optionalRequest.Ts != "" {
		request.AddParam("ts", optionalRequest.Ts)
	}

	url := p.publicUrlBuilder.Build("/v1/settings/common/market-symbols", request)

	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := common.GetMarketSymbolsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)

	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// 查询链信息
func (p *CommonClient) GetChains(optionalRequest common.GetChains) ([]common.ChainsV1, error) {
	request := new(model.GetRequest).Init()
	if optionalRequest.ShowDesc != "" {
		request.AddParam("show-desc", optionalRequest.ShowDesc)
	}
	if optionalRequest.Currency != "" {
		request.AddParam("currency", optionalRequest.Currency)
	}
	if optionalRequest.Ts != "" {
		request.AddParam("ts", optionalRequest.Ts)
	}

	url := p.publicUrlBuilder.Build("/v1/settings/common/chains", request)

	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := common.GetChainsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)

	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}

	return nil, errors.New(getResp)
}
