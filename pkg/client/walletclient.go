package client

import (
	"../../internal"
	"../../internal/requestbuilder"
	"../../pkg/getrequest"
	"../../pkg/postrequest"
	"../../pkg/response/wallet"
	"encoding/json"
	"errors"
	"strconv"
)

type WalletClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

func (p *WalletClient) Init(accessKey string, secretKey string, host string) *WalletClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return p
}

func (p *WalletClient) GetDepositAddress(currency string) ([]wallet.DepositAddress, error) {
	request := new(getrequest.GetRequest).Init()

	request.AddParam("currency", currency)

	url := p.privateUrlBuilder.Build("GET", "/v2/account/deposit/address", request)
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

func (p *WalletClient) GetWithdrawQuota(currency string) (*wallet.WithdrawQuota, error) {
	request := new(getrequest.GetRequest).Init()

	request.AddParam("currency", currency)

	url := p.privateUrlBuilder.Build("GET", "/v2/account/withdraw/quota", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := wallet.GetWithdrawQuotaResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}
func (p *WalletClient) CreateWithdraw(request postrequest.CreateWithdrawRequest) (int64, error) {
	postBody, jsonErr := postrequest.ToJson(request)

	url := p.privateUrlBuilder.Build("POST", "/v1/dw/withdraw/api/create", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return 0, postErr
	}

	result := wallet.CreateWithdrawResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return 0, jsonErr
	}

	if result.Status == "ok" && result.Data != 0 {
		return result.Data, nil
	}
	return 0, errors.New(postResp)

}
func (p *WalletClient) CancelWithdraw(withdrawId int64) (int64, error) {

	url := p.privateUrlBuilder.Build("POST", "/v1/dw/withdraw-virtual/"+strconv.FormatInt(withdrawId, 10)+"}/cancel", nil)
	postResp, postErr := internal.HttpPost(url, "")
	if postErr != nil {
		return 0, postErr
	}
	result := wallet.CancelWithdrawResponse{}
	jsonErr := json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return 0, jsonErr
	}

	if result.Status == "ok" && result.Data != 0 {
		return result.Data, nil
	}
	return 0, errors.New(postResp)

}

func (p *WalletClient) QueryDepositWithdraw(depositOrWithdraw string, optionalRequest getrequest.QueryDepositWithdrawOptionalRequest) ([]wallet.DepositWithdraw, error) {
	request := new(getrequest.GetRequest).Init()

	request.AddParam("type", depositOrWithdraw)

	if optionalRequest.Currency != "" {
		request.AddParam("currency", optionalRequest.Currency)
	}
	if optionalRequest.From != "" {
		request.AddParam("from", optionalRequest.From)
	}
	if optionalRequest.Direct != "" {
		request.AddParam("direct", optionalRequest.Direct)
	}
	if optionalRequest.Size != "" {
		request.AddParam("size", optionalRequest.Size)
	}

	url := p.privateUrlBuilder.Build("GET", "/v1/query/deposit-withdraw", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := wallet.QueryDepositWithdrawResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}
