package client

import (
	"encoding/json"
	"errors"
	"github.com/huobirdcenter/huobi_golang/internal"
	"github.com/huobirdcenter/huobi_golang/internal/requestbuilder"
	"github.com/huobirdcenter/huobi_golang/pkg/getrequest"
	"github.com/huobirdcenter/huobi_golang/pkg/model/subuser"
	"github.com/huobirdcenter/huobi_golang/pkg/model/wallet"
	"github.com/huobirdcenter/huobi_golang/pkg/postrequest"
	"strconv"
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
	postBody, jsonErr := postrequest.ToJson(request)

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

// Parent user query sub user deposit address of corresponding chain, for a specific crypto currency (except IOTA)
func (p *SubUserClient) GetSubUserDepositAddress(subUid int64, currency string) ([]wallet.DepositAddress, error) {
	request := new(getrequest.GetRequest).Init()
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
	request := new(getrequest.GetRequest).Init()

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