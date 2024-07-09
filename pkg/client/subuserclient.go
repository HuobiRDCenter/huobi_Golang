package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/internal"
	"github.com/huobirdcenter/huobi_golang/internal/requestbuilder"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/account"
	"github.com/huobirdcenter/huobi_golang/pkg/model/subuser"
	"github.com/huobirdcenter/huobi_golang/pkg/model/wallet"
	"strconv"
	"strings"
)

// Responsible to operate wallet
type SubUserClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

// Initializer
func (p *SubUserClient) Init(accessKey string, secretKey string, host string) *SubUserClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return p
}

// Parent user query sub user deposit address of corresponding chain, for a specific crypto currency (except IOTA)
func (p *SubUserClient) CreateSubUser(request subuser.CreateSubUserRequest) ([]subuser.UserData, error) {
	postBody, jsonErr := model.ToJson(request)

	url := p.privateUrlBuilder.Build("POST", "/v2/sub-user/creation", nil)
	postResp, postErr := internal.HttpPost(url, string(postBody))
	if postErr != nil {
		return nil, postErr
	}

	result := subuser.CreateSubUserResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(postResp)
}

// Lock or unlock a specific user
func (p *SubUserClient) SubUserManagement(request subuser.SubUserManagementRequest) (*subuser.SubUserManagement, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/sub-user/management", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := subuser.SubUserManagementResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code != 200 {
		return nil, errors.New(postResp)
	}
	return result.Data, nil
}

// Set Tradable Market for Sub Users
func (p *SubUserClient) SetSubUserTradableMarket(request subuser.SetSubUserTradableMarketRequest) ([]subuser.TradableMarket, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/sub-user/tradable-market", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := subuser.SetSubUserTradableMarketResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code != 200 {
		return nil, errors.New(postResp)
	}

	return result.Data, nil
}

// Set Asset Transfer Permission for Sub Users
func (p *SubUserClient) SetSubUserTransferability(request subuser.SetSubUserTransferabilityRequest) ([]subuser.Transferability, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/sub-user/transferability", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := subuser.SetSubUserTransferabilityResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code != 200 {
		return nil, errors.New(postResp)
	}

	return result.Data, nil
}

// Transfer asset between parent and sub account
func (p *SubUserClient) SubUserTransfer(request subuser.SubUserTransferRequest) (string, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return "", jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/subuser/transfer", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return "", postErr
	}
	if strings.Contains(postResp, "data") {
		return postResp, nil
	} else {
		return "", errors.New(postResp)
	}
}

// Parent user query sub user deposit address of corresponding chain, for a specific crypto currency (except IOTA)
func (p *SubUserClient) GetSubUserDepositAddress(subUid int64, currency string) ([]wallet.DepositAddress, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("subUid", strconv.FormatInt(subUid, 10))
	request.AddParam("currency", currency)

	url := p.privateUrlBuilder.Build("GET", "/v2/sub-user/deposit-address", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := wallet.GetDepositAddressResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}

// Parent user query sub user deposits history
func (p *SubUserClient) QuerySubUserDepositHistory(subUid int64, optionalRequest subuser.QuerySubUserDepositHistoryOptionalRequest) ([]subuser.DepositHistory, error) {
	request := new(model.GetRequest).Init()

	request.AddParam("subUid", strconv.FormatInt(subUid, 10))

	if optionalRequest.Currency != "" {
		request.AddParam("currency", optionalRequest.Currency)
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
	if optionalRequest.Limit != "" {
		request.AddParam("limit", optionalRequest.Limit)
	}
	if optionalRequest.FromId != 0 {
		request.AddParam("fromId", strconv.FormatInt(optionalRequest.FromId, 10))
	}

	url := p.privateUrlBuilder.Build("GET", "/v2/sub-user/query-deposit", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := subuser.QuerySubUserDepositHistoryResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}

// Returns the aggregated balance from all the sub-users
func (p *SubUserClient) GetSubUserAggregateBalance() ([]account.AggregateBalance, error) {
	url := p.privateUrlBuilder.Build("GET", "/v1/subuser/aggregate-balance", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := account.GetSubUserAggregateBalanceResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// Returns the balance of a sub-account specified by sub-uid
func (p *SubUserClient) GetSubUserAccount(subUid int64) ([]account.SubUserAccount, error) {
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

func (p *SubUserClient) GetUid() (int64, error) {
	url := p.privateUrlBuilder.Build("GET", "/v2/user/uid", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return 0, getErr
	}
	result := account.GetUidResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return 0, jsonErr
	}
	if result.Code == 200 && result.Data != 0 {
		return result.Data, nil
	}
	return 0, errors.New(getResp)
}

// 设置子用户手续费抵扣模式
func (p *SubUserClient) DeductMode(request subuser.DeductModeRequest) (*subuser.DeductModeResponse, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/sub-user/deduct-mode", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := subuser.DeductModeResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}

	return nil, errors.New(postResp)
}

// 母子用户APIkey信息查询
func (p *SubUserClient) GetApiKey(uid int64, optionalRequest subuser.GetApiKey) ([]subuser.ApiKey, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("uid", strconv.FormatInt(uid, 10))
	if optionalRequest.AccessKey != "" {
		request.AddParam("accessKey", optionalRequest.AccessKey)
	}

	url := p.privateUrlBuilder.Build("GET", "/v2/user/api-key", request)

	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := subuser.GetApiKeyResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)

	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// 获取子用户列表
func (p *SubUserClient) GetUserList(optionalRequest subuser.GetUserList) ([]subuser.UserList, error) {
	request := new(model.GetRequest).Init()
	if optionalRequest.FromId != 0 {
		request.AddParam("fromId", strconv.FormatInt(optionalRequest.FromId, 10))
	}

	url := p.privateUrlBuilder.Build("GET", "/v2/sub-user/user-list", request)

	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := subuser.GetUserListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)

	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// 获取特定子用户的用户状态
func (p *SubUserClient) GetUserState(subUid int64) (*subuser.GetUserStateResponse, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("subUid", strconv.FormatInt(subUid, 10))

	url := p.privateUrlBuilder.Build("GET", "/v2/sub-user/user-state", request)

	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := subuser.GetUserStateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}

	return nil, errors.New(getResp)
}

// 获取特定子用户的账户列表
func (p *SubUserClient) GetAccountList(subUid int64) (*subuser.GetAccountListResponse, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("subUid", strconv.FormatInt(subUid, 10))

	url := p.privateUrlBuilder.Build("GET", "/v2/sub-user/account-list", request)

	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := subuser.GetAccountListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}

	return nil, errors.New(getResp)
}

// 子用户APIkey创建
func (p *SubUserClient) ApiKeyGeneration(request subuser.ApiKeyGenerationRequest) (*subuser.ApiKeyGenerationResponse, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/sub-user/api-key-generation", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := subuser.ApiKeyGenerationResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}

	return nil, errors.New(postResp)
}

// 修改子用户APIkey
func (p *SubUserClient) ApiKeyModification(request subuser.ApiKeyModificationRequest) (*subuser.ApiKeyModificationResponse, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/sub-user/api-key-modification", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := subuser.ApiKeyModificationResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}

	return nil, errors.New(postResp)
}

// 删除子用户APIkey
func (p *SubUserClient) ApiKeyDeletion(request subuser.ApiKeyDeletionRequest) (*subuser.ApiKeyDeletionResponse, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/sub-user/api-key-deletion", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := subuser.ApiKeyDeletionResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}

	return nil, errors.New(postResp)
}

// 用户主动授信
func (p *SubUserClient) ActiveCredit(request subuser.ActiveCreditRequest) (bool, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return false, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/trust/user/active/credit", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return false, postErr
	}

	result := subuser.ActiveCreditResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return false, jsonErr
	}
	if result.Status == "ok" {
		return true, nil
	}

	return false, errors.New(postResp)
}
