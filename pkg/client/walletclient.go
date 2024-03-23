package client

import (
	"encoding/json"
	"errors"
	"github.com/huobirdcenter/huobi_golang/internal"
	"github.com/huobirdcenter/huobi_golang/internal/requestbuilder"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/wallet"
	"strconv"
)

// Responsible to operate wallet
type WalletClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

// Initializer
func (p *WalletClient) Init(accessKey string, secretKey string, host string) *WalletClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return p
}

// Get deposit address of corresponding chain, for a specific crypto currency (except IOTA)
func (p *WalletClient) GetDepositAddress(currency string) ([]wallet.DepositAddress, error) {
	request := new(model.GetRequest).Init()

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

// Query withdraw quota for currencies
func (p *WalletClient) GetWithdrawQuota(currency string) (*wallet.WithdrawQuota, error) {
	request := new(model.GetRequest).Init()

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

// Parent user to query withdraw address available for API key
func (p *WalletClient) GetWithdrawAddress(request *model.GetRequest) (*wallet.GetWithdrawAddressResponse, error) {
	url := p.privateUrlBuilder.Build("GET", "/v2/account/withdraw/address", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := wallet.GetWithdrawAddressResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}
	return nil, errors.New(getResp)
}

// Withdraw from spot trading account to an external address.
func (p *WalletClient) CreateWithdraw(request wallet.CreateWithdrawRequest) (int64, error) {
	postBody, jsonErr := model.ToJson(request)

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

// Cancels a previously created withdraw request by its transfer id.
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

// Returns all existed withdraws and deposits and return their latest status.
func (p *WalletClient) QueryDepositWithdraw(depositOrWithdraw string, optionalRequest wallet.QueryDepositWithdrawOptionalRequest) ([]wallet.DepositWithdraw, error) {
	request := new(model.GetRequest).Init()

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

// 通过clientOrderId查询提币订单
func (p *WalletClient) GetWithdraw(clientOrderId string) (*wallet.GetWithdrawResponse, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("clientOrderId", clientOrderId)

	url := p.privateUrlBuilder.Build("GET", "/v1/query/withdraw/client-order-id", request)

	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := wallet.GetWithdrawResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" {
		return &result, nil
	}

	return nil, errors.New(getResp)
}
