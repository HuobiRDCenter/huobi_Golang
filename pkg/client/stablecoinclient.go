package client

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/huobirdcenter/huobi_golang/internal"
	"github.com/huobirdcenter/huobi_golang/internal/requestbuilder"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/stablecoin"
)

// Responsible to operate wallet
type StableCoinClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

// Initializer
func (p *StableCoinClient) Init(accessKey string, secretKey string, host string, s string) *StableCoinClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host, s)
	return p
}

// Get stable coin exchange rate
func (p *StableCoinClient) GetExchangeRate(currency string, amount string, exchangeType string) (*stablecoin.GetExchangeRateResponse, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("currency", currency)
	request.AddParam("amount", amount)
	request.AddParam("type", exchangeType)

	url := p.privateUrlBuilder.Build("GET", "/v1/stable-coin/quote", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := stablecoin.GetExchangeRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {
		return &result, nil
	}
	return nil, errors.New(getResp)
}

// Exchange stable coin
func (p *StableCoinClient) ExchangeStableCoin(quoteId string) (*stablecoin.ExchangeStableCoinResponse, error) {
	postBody := fmt.Sprintf("{ \"quote-id\": \"%s\"}", quoteId)

	url := p.privateUrlBuilder.Build("POST", "/v1/stable-coin/exchange", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := stablecoin.ExchangeStableCoinResponse{}
	jsonErr := json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Status == "ok" && result.Data != nil {
		return &result, nil
	}
	return nil, errors.New(postResp)
}
