package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/internal"
	"github.com/huobirdcenter/huobi_golang/internal/requestbuilder"
	"github.com/huobirdcenter/huobi_golang/pkg/model/account"
	"github.com/huobirdcenter/huobi_golang/pkg/util"
	"strconv"
)

// Responsible to operate account
type AccountClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

// Initializer
func (p *AccountClient) Init(accessKey string, secretKey string, host string) *AccountClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
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

func (p *AccountClient) TransferAccount(request account.TransferAccountRequest) (*account.TransferAccountResponse, error) {
	postBody, jsonErr := util.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/subuser/transfer", nil)
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
	request := new(util.GetRequest).Init()
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
	request := new(util.GetRequest).Init()
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
	postBody, jsonErr := util.ToJson(request)
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

// Returns the balance of a sub-account specified by sub-uid
func (p *AccountClient) GetSubUserAccount(subUid int64) ([]account.SubUserAccount, error) {
	url := p.privateUrlBuilder.Build("GET", fmt.Sprintf("/v1/account/accounts/%d", subUid), nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := account.GetSubUserAccountResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}

	return nil, errors.New(getResp)
}
