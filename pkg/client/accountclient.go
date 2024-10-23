package client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/huobirdcenter/huobi_golang/internal"
	"github.com/huobirdcenter/huobi_golang/internal/requestbuilder"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/account"
)

// Responsible to operate account
type AccountClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

// Initializer
func (p *AccountClient) Init(accessKey string, secretKey string, host string, sign string) *AccountClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host, sign)
	return p
}

// Returns a list of accounts owned by this API user
func (p *AccountClient) GetAccountInfo() ([]account.AccountInfo, error) {
	url := p.privateUrlBuilder.Build("GET", "/v1/account/accounts", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := account.GetAccountInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// Returns the balance of an account specified by account id
func (p *AccountClient) GetAccountBalance(accountId string) (*account.AccountBalance, error) {
	url := p.privateUrlBuilder.Build("GET", "/v1/account/accounts/"+accountId+"/balance", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := account.GetAccountBalanceResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// Returns the valuation of the total assets of the account in btc or fiat currency.
func (p *AccountClient) GetAccountAssetValuation(accountType string, valuationCurrency string, subUid int64) (*account.GetAccountAssetValuationResponse, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("accountType", accountType)
	if valuationCurrency != "" {
		request.AddParam("valuationCurrency", valuationCurrency)
	}
	if subUid != 0 {
		request.AddParam("subUid", strconv.FormatInt(subUid, 10))
	}

	url := p.privateUrlBuilder.Build("GET", "/v2/account/asset-valuation", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := account.GetAccountAssetValuationResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}

	return nil, errors.New(getResp)
}

func (p *AccountClient) TransferAccount(request account.TransferAccountRequest) (*account.TransferAccountResponse, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/account/transfer", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := account.TransferAccountResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status != "ok" {
		return nil, errors.New(postResp)
	}

	return &result, nil
}

// Returns the amount changes of specified user's account
func (p *AccountClient) GetAccountHistory(accountId string, optionalRequest account.GetAccountHistoryOptionalRequest) ([]account.AccountHistory, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("account-id", accountId)
	if optionalRequest.Currency != "" {
		request.AddParam("currency", optionalRequest.Currency)
	}
	if optionalRequest.Size != 0 {
		request.AddParam("size", strconv.Itoa(optionalRequest.Size))
	}
	if optionalRequest.EndTime != 0 {
		request.AddParam("end-time", strconv.FormatInt(optionalRequest.EndTime, 10))
	}
	if optionalRequest.Sort != "" {
		request.AddParam("sort", optionalRequest.Sort)
	}
	if optionalRequest.StartTime != 0 {
		request.AddParam("start-time", strconv.FormatInt(optionalRequest.StartTime, 10))
	}
	if optionalRequest.TransactTypes != "" {
		request.AddParam("transact-types", optionalRequest.TransactTypes)
	}

	url := p.privateUrlBuilder.Build("GET", "/v1/account/history", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := account.GetAccountHistoryResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// Returns the account ledger of specified user's account
func (p *AccountClient) GetAccountLedger(accountId string, optionalRequest account.GetAccountLedgerOptionalRequest) ([]account.Ledger, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("accountId", accountId)
	if optionalRequest.Currency != "" {
		request.AddParam("currency", optionalRequest.Currency)
	}
	if optionalRequest.TransactTypes != "" {
		request.AddParam("transactTypes", optionalRequest.TransactTypes)
	}
	if optionalRequest.StartTime != 0 {
		request.AddParam("startTime", strconv.FormatInt(optionalRequest.StartTime, 10))
	}
	if optionalRequest.EndTime != 0 {
		request.AddParam("endTime", strconv.FormatInt(optionalRequest.EndTime, 10))
	}
	if optionalRequest.Sort != "" {
		request.AddParam("sort", optionalRequest.Sort)
	}
	if optionalRequest.Limit != 0 {
		request.AddParam("limit", strconv.Itoa(optionalRequest.Limit))
	}
	if optionalRequest.FromId != 0 {
		request.AddParam("limit", strconv.FormatInt(optionalRequest.EndTime, 10))
	}

	url := p.privateUrlBuilder.Build("GET", "/v2/account/ledger", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := account.GetAccountLedgerResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// Transfer fund between spot account and future contract account
func (p *AccountClient) FuturesTransfer(request account.FuturesTransferRequest) (int64, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return 0, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/futures/transfer", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return 0, postErr
	}

	result := account.FuturesTransferResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return 0, jsonErr
	}
	if result.Status != "ok" {
		return 0, errors.New(postResp)

	}
	return result.Data, nil
}

// Returns the point balance of specified user's account
func (p *AccountClient) GetPointBalance(subUid string) (*account.GetPointBalanceResponse, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("subUid", subUid)

	url := p.privateUrlBuilder.Build("GET", "/v2/point/account", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := account.GetPointBalanceResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}

	return nil, errors.New(getResp)
}

// Transfer points between spot account and future contract account
func (p *AccountClient) TransferPoint(request account.TransferPointRequest) (*account.TransferPointResponse, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/point/transfer", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := account.TransferPointResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}

	return nil, errors.New(postResp)
}

// 获取平台资产总估值
func (p *AccountClient) GetValuation(accountType string, optionalRequest account.GetValuation) (*account.GetValuationResponse, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("accountType", accountType)
	if optionalRequest.ValuationCurrency != "" {
		request.AddParam("valuationCurrency", optionalRequest.ValuationCurrency)
	}

	url := p.privateUrlBuilder.Build("GET", "/v2/account/valuation", request)

	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := account.GetValuationResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}

	return nil, errors.New(getResp)
}

// 【通用】现货-合约账户和OTC账户间进行资金的划转
func (p *AccountClient) Transfer(request account.TransferRequest) (*account.TransferResponse, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/account/transfer", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := account.TransferResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}

	return nil, errors.New(postResp)
}

// 用户抵扣信息查询
func (p *AccountClient) GetUserInfo() (*account.GetUserInfoResponse, error) {
	url := p.privateUrlBuilder.Build("GET", "/v1/account/switch/user/info", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := account.GetUserInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}

	return nil, errors.New(getResp)
}

// 可抵扣币种查询信息
func (p *AccountClient) GetOverviewInfo() (*account.GetOverviewInfoResponse, error) {
	url := p.privateUrlBuilder.Build("GET", "/v1/account/overview/info", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := account.GetOverviewInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}

	return nil, errors.New(getResp)
}

// 设置现货/杠杆抵扣手续费方式
func (p *AccountClient) FeeSwitch(request account.FeeSwitchRequest) (*account.FeeSwitchResponse, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/account/fee/switch", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := account.FeeSwitchResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}

	return nil, errors.New(postResp)
}
