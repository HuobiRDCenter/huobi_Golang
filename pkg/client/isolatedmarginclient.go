package client

import (
	"../../internal"
	"../../internal/requestbuilder"
	"../../pkg/getrequest"
	"../../pkg/postrequest"
	"../response/margin"
	"encoding/json"
	"errors"
	"strconv"
)

type IsolatedMarginClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

func (p *IsolatedMarginClient) Init(accessKey string, secretKey string, host string) *IsolatedMarginClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return p
}

func (p *IsolatedMarginClient) TransferIn(request postrequest.IsolatedMarginTransferRequest) (int, error) {

	postBody, jsonErr := postrequest.ToJson(request)
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

func (p *IsolatedMarginClient) TransferOut(request postrequest.IsolatedMarginTransferRequest) (int, error) {

	postBody, jsonErr := postrequest.ToJson(request)
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

func (p *IsolatedMarginClient) GetMarginLoanInfo(optionalRequest getrequest.GetMarginLoanInfoOptionalRequest) ([]margin.IsolatedMarginLoanInfo, error) {
	request := new(getrequest.GetRequest).Init()
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
func (p *IsolatedMarginClient) MarginOrders(request postrequest.IsolatedMarginOrdersRequest) (int, error) {
	postBody, jsonErr := postrequest.ToJson(request)
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
func (p *IsolatedMarginClient) MarginOrdersRepay(orderId string, request postrequest.MarginOrdersRepayRequest) (int, error) {
	postBody, jsonErr := postrequest.ToJson(request)
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

func (p *IsolatedMarginClient) MarginLoanOrders(symbol string, optionalRequest getrequest.IsolatedMarginLoanOrdersOptionalRequest) ([]margin.IsolatedMarginLoanOrder, error) {
	request := new(getrequest.GetRequest).Init()
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
func (p *IsolatedMarginClient) MarginAccountsBalance(optionalRequest getrequest.MarginAccountsBalanceOptionalRequest) ([]margin.IsolatedMarginAccountsBalance, error) {

	request := new(getrequest.GetRequest).Init()
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
