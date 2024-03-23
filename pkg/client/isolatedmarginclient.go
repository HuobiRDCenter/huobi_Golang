package client

import (
	"encoding/json"
	"errors"
	"github.com/huobirdcenter/huobi_golang/internal"
	"github.com/huobirdcenter/huobi_golang/internal/requestbuilder"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/margin"
	"strconv"
)

// Responsible to operate isolated margin
type IsolatedMarginClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

// Initializer
func (p *IsolatedMarginClient) Init(accessKey string, secretKey string, host string) *IsolatedMarginClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return p
}

// Transfer specific asset from spot trading account to isolated margin account
func (p *IsolatedMarginClient) TransferIn(request margin.IsolatedMarginTransferRequest) (int, error) {

	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return 0, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/dw/transfer-in/margin", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return 0, postErr
	}

	result := margin.TransferResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return 0, jsonErr
	}
	if result.Status != "ok" {
		return 0, errors.New(postResp)

	}
	return result.Data, nil
}

// Transfer specific asset from isolated margin account to spot trading account
func (p *IsolatedMarginClient) TransferOut(request margin.IsolatedMarginTransferRequest) (int, error) {

	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return 0, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/dw/transfer-out/margin", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return 0, postErr
	}

	result := margin.TransferResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return 0, jsonErr
	}
	if result.Status != "ok" {
		return 0, errors.New(postResp)
	}
	return result.Data, nil
}

// Returns loan interest rates and quota applied on the user
func (p *IsolatedMarginClient) GetMarginLoanInfo(optionalRequest margin.GetMarginLoanInfoOptionalRequest) ([]margin.IsolatedMarginLoanInfo, error) {
	request := new(model.GetRequest).Init()
	if optionalRequest.Symbols != "" {
		request.AddParam("symbols", optionalRequest.Symbols)
	}
	url := p.privateUrlBuilder.Build("GET", "/v1/margin/loan-info", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := margin.GetIsolatedMarginLoanInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// Place an order to apply a margin loan.
func (p *IsolatedMarginClient) Apply(request margin.IsolatedMarginOrdersRequest) (int, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return 0, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/margin/orders", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return 0, postErr
	}

	result := margin.MarginOrdersResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return 0, jsonErr
	}
	if result.Status != "ok" {
		return 0, errors.New(postResp)

	}
	return result.Data, nil

}

// Repays margin loan with you asset in your margin account.
func (p *IsolatedMarginClient) Repay(orderId string, request margin.MarginOrdersRepayRequest) (int, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return 0, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/margin/orders/"+orderId+"/repay", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return 0, postErr
	}

	result := margin.MarginOrdersRepayResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return 0, jsonErr
	}
	if result.Status != "ok" {
		return 0, errors.New(postResp)

	}
	return result.Data, nil
}

// Returns margin orders based on a specific searching criteria.
func (p *IsolatedMarginClient) MarginLoanOrders(symbol string, optionalRequest margin.IsolatedMarginLoanOrdersOptionalRequest) ([]margin.IsolatedMarginLoanOrder, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("symbol", symbol)
	if optionalRequest.Size != "" {
		request.AddParam("size", optionalRequest.Size)
	}
	if optionalRequest.Direct != "" {
		request.AddParam("direct", optionalRequest.Direct)
	}
	if optionalRequest.EndDate != "" {
		request.AddParam("end-date", optionalRequest.EndDate)
	}
	if optionalRequest.From != "" {
		request.AddParam("from", optionalRequest.From)
	}
	if optionalRequest.StartDate != "" {
		request.AddParam("start-date", optionalRequest.StartDate)
	}
	if optionalRequest.States != "" {
		request.AddParam("states", optionalRequest.States)
	}
	if optionalRequest.SubUid != 0 {
		request.AddParam("sub-uid", strconv.Itoa(optionalRequest.SubUid))
	}
	url := p.privateUrlBuilder.Build("GET", "/v1/margin/loan-orders", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := margin.IsolatedMarginLoanOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// Returns the balance of the margin loan account.
func (p *IsolatedMarginClient) MarginAccountsBalance(optionalRequest margin.MarginAccountsBalanceOptionalRequest) ([]margin.IsolatedMarginAccountsBalance, error) {

	request := new(model.GetRequest).Init()
	if optionalRequest.SubUid != 0 {
		request.AddParam("sub-uid", strconv.Itoa(optionalRequest.SubUid))
	}
	if optionalRequest.Symbol != "" {
		request.AddParam("symbol", optionalRequest.Symbol)
	}
	url := p.privateUrlBuilder.Build("GET", "/v1/margin/accounts/balance", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := margin.IsolatedMarginAccountsBalanceResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// 获取杠杆持仓限额（全仓）
func (p *IsolatedMarginClient) GetMarginLimit(currency string) ([]margin.MarginLimit, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("currency", currency)

	url := p.privateUrlBuilder.Build("GET", "/v2/margin/limit", request)

	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := margin.GetMarginLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)

	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}

	return nil, errors.New(getResp)
}
