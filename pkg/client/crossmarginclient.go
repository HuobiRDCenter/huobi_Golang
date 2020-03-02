package client

import (
	"../../internal"
	"../../internal/requestbuilder"
	"../../pkg/getrequest"
	"../../pkg/postrequest"
	"../response/margin"
	"encoding/json"
	"errors"
)

type CrossMarginClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

func (p *CrossMarginClient) Init(accessKey string, secretKey string, host string) *CrossMarginClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return p
}

func (p *CrossMarginClient) TransferIn(request postrequest.CrossMarginTransferRequest) (int, error) {

	postBody, jsonErr := postrequest.ToJson(request)
	if jsonErr != nil {
		return 0, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/cross-margin/transfer-in", nil)
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

func (p *CrossMarginClient) TransferOut(request postrequest.CrossMarginTransferRequest) (int, error) {

	postBody, jsonErr := postrequest.ToJson(request)
	if jsonErr != nil {
		return 0, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/cross-margin/transfer-out", nil)
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

func (p *CrossMarginClient) GetMarginLoanInfo() ([]margin.CrossMarginLoanInfo, error) {
	request := new(getrequest.GetRequest).Init()

	url := p.privateUrlBuilder.Build("GET", "/v1/cross-margin/loan-info", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := margin.GetCrossMarginLoanInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}

	return nil, errors.New(getResp)

}
func (p *CrossMarginClient) MarginOrders(request postrequest.CrossMarginOrdersRequest) (int, error) {
	postBody, jsonErr := postrequest.ToJson(request)
	if jsonErr != nil {
		return 0, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/cross-margin/orders", nil)
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
func (p *CrossMarginClient) MarginOrdersRepay(orderId string, request postrequest.MarginOrdersRepayRequest) error {
	postBody, jsonErr := postrequest.ToJson(request)
	if jsonErr != nil {
		return jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/cross-margin/orders/"+orderId+"/repay", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return postErr
	}

	result := margin.MarginOrdersRepayResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return jsonErr
	}
	if result.Status != "ok" {
		return errors.New(postResp)

	}
	return nil
}

func (p *CrossMarginClient) MarginLoanOrders(optionalRequest getrequest.CrossMarginLoanOrdersOptionalRequest) ([]margin.CrossMarginLoanOrder, error) {
	request := new(getrequest.GetRequest).Init()
	if optionalRequest.Size != "" {
		request.AddParam("size", optionalRequest.Size)
	}
	if optionalRequest.Currency != "" {
		request.AddParam("currency", optionalRequest.Currency)

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
	if optionalRequest.State != "" {
		request.AddParam("state", optionalRequest.State)
	}

	url := p.privateUrlBuilder.Build("GET", "/v1/cross-margin/loan-orders", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := margin.CrossMarginLoanOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}

	return nil, errors.New(getResp)
}
func (p *CrossMarginClient) MarginAccountsBalance() (*margin.CrossMarginAccountsBalance, error) {

	request := new(getrequest.GetRequest).Init()
	url := p.privateUrlBuilder.Build("GET", "/v1/cross-margin/accounts/balance", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := margin.CrossMarginAccountsBalanceResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}

	return nil, errors.New(getResp)
}
