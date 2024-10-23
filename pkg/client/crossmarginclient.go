package client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/huobirdcenter/huobi_golang/internal"
	"github.com/huobirdcenter/huobi_golang/internal/requestbuilder"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/margin"
)

// Responsible to operate cross margin
type CrossMarginClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

// Initializer
func (p *CrossMarginClient) Init(accessKey string, secretKey string, host string, s string) *CrossMarginClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host, s)
	return p
}

// Transfer specific asset from spot trading account to cross margin account
func (p *CrossMarginClient) TransferIn(request margin.CrossMarginTransferRequest) (int, error) {

	postBody, jsonErr := model.ToJson(request)
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

// Transfer specific asset from cross margin account to spot trading account
func (p *CrossMarginClient) TransferOut(request margin.CrossMarginTransferRequest) (int, error) {

	postBody, jsonErr := model.ToJson(request)
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

// Returns loan interest rates and quota applied on the user
func (p *CrossMarginClient) GetMarginLoanInfo() ([]margin.CrossMarginLoanInfo, error) {
	request := new(model.GetRequest).Init()

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

// Place an order to apply a margin loan.
func (p *CrossMarginClient) ApplyLoan(request margin.CrossMarginOrdersRequest) (int, error) {
	postBody, jsonErr := model.ToJson(request)
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

// Repays margin loan with you asset in your margin account.
func (p *CrossMarginClient) Repay(orderId string, request margin.MarginOrdersRepayRequest) (int, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return 0, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/cross-margin/orders/"+orderId+"/repay", nil)
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
func (p *CrossMarginClient) MarginLoanOrders(optionalRequest margin.CrossMarginLoanOrdersOptionalRequest) ([]margin.CrossMarginLoanOrder, error) {
	request := new(model.GetRequest).Init()
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
	if optionalRequest.SubUid != "" {
		request.AddParam("sub-uid", optionalRequest.SubUid)
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

// Returns the balance of the margin loan account.
func (p *CrossMarginClient) MarginAccountsBalance(SubUid string) (*margin.CrossMarginAccountsBalance, error) {

	request := new(model.GetRequest).Init()
	request.AddParam("sub-uid", SubUid)

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

// Repays general margin loan with you asset in your margin account.
func (p *CrossMarginClient) GeneralRepay(request margin.CrossMarginGeneralReplayLoanOptionalRequest) ([]margin.CrossMarginGeneraReplaylLoan, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/account/repayment", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := margin.CrossMarginGeneralReplyLoanResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code != 200 {
		return nil, errors.New(postResp)

	}
	return result.Data, nil
}

// Returns general margin orders based on a specific searching criteria.
func (p *CrossMarginClient) GeneralMarginLoanOrders(optionalRequest margin.CrossMarginGeneralReplayLoanRecordsOptionalRequest) ([]margin.CrossMarginGeneraReplaylLoanRecord, error) {
	request := new(model.GetRequest).Init()
	if optionalRequest.RepayId != "" {
		request.AddParam("repayId", optionalRequest.RepayId)
	}
	if optionalRequest.AccountId != "" {
		request.AddParam("accountId", optionalRequest.AccountId)
	}
	if optionalRequest.Currency != "" {
		request.AddParam("currency", optionalRequest.Currency)
	}
	if optionalRequest.StartDate != 0 {
		request.AddParam("startDate", strconv.FormatInt(optionalRequest.StartDate, 10))
	}
	if optionalRequest.EndDate != 0 {
		request.AddParam("endDate", strconv.FormatInt(optionalRequest.EndDate, 10))
	}
	if optionalRequest.Sort != "" {
		request.AddParam("sort", optionalRequest.Sort)
	}
	if optionalRequest.Limit != 0 {
		request.AddParam("limit", strconv.Itoa(optionalRequest.Limit))
	}
	if optionalRequest.FromId != 0 {
		request.AddParam("fromId", strconv.FormatInt(optionalRequest.FromId, 10))
	}

	url := p.privateUrlBuilder.Build("GET", "/v2/account/repayment", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := margin.CrossMarginGeneralReplyLoanRecordsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}

	return nil, errors.New(getResp)
}
