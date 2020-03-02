package client

import (
	"../../internal"
	"../../internal/requestbuilder"
	"../../pkg/getrequest"
	"../../pkg/postrequest"
	"../../pkg/response/etf"
	"encoding/json"
	"errors"
	"strconv"
)

type ETFClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

func (p *ETFClient) Init(accessKey string, secretKey string, host string) *ETFClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return p
}

func (p *ETFClient) GetSwapConfig(etfName string) (*etf.SwapConfig, error) {
	request := new(getrequest.GetRequest).Init()

	request.AddParam("etf_name", etfName)

	url := p.privateUrlBuilder.Build("GET", "/etf/swap/config", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := etf.GetSwapConfigResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}
func (p *ETFClient) SwapIn(request postrequest.SwapRequest) (bool, error) {

	postBody, jsonErr := postrequest.ToJson(request)
	if jsonErr != nil {
		return false, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/etf/swap/in", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return false, postErr
	}
	result := etf.SwapResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return false, jsonErr
	}

	if result.Code == 200 && result.Success == true {
		return result.Success, nil
	}
	return false, errors.New(postResp)

}

func (p *ETFClient) SwapOut(request postrequest.SwapRequest) (bool, error) {

	postBody, jsonErr := postrequest.ToJson(request)
	if jsonErr != nil {
		return false, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/etf/swap/out", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return false, postErr
	}
	result := etf.SwapResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return false, jsonErr
	}

	if result.Code == 200 && result.Success == true {
		return result.Success, nil
	}
	return false, errors.New(postResp)

}
func (p *ETFClient) GetSwapList(etfName string, offset int, limit int) ([]*etf.SwapList, error) {
	request := new(getrequest.GetRequest).Init()

	request.AddParam("etf_name", etfName)
	request.AddParam("offset", strconv.Itoa(offset))
	request.AddParam("limit", strconv.Itoa(limit))

	url := p.privateUrlBuilder.Build("GET", "/etf/swap/list", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := etf.GetSwapListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}
