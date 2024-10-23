package client

import (
	"encoding/json"
	"errors"

	"github.com/huobirdcenter/huobi_golang/internal"
	"github.com/huobirdcenter/huobi_golang/internal/requestbuilder"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/algoorder"
)

// Responsible to operate algo order
type AlgoOrderClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

// Initializer
func (p *AlgoOrderClient) Init(accessKey string, secretKey string, host string, s string) *AlgoOrderClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host, s)
	return p
}

// Place a new order
func (p *AlgoOrderClient) PlaceOrder(request *algoorder.PlaceOrderRequest) (*algoorder.PlaceOrderResponse, error) {
	postBody, jsonErr := model.ToJson(request)

	url := p.privateUrlBuilder.Build("POST", "/v2/algo-orders", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := algoorder.PlaceOrderResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

// Cancel orders by client order id
func (p *AlgoOrderClient) CancelOrder(request *algoorder.CancelOrdersRequest) (*algoorder.CancelOrdersResponse, error) {
	postBody, jsonErr := model.ToJson(request)

	url := p.privateUrlBuilder.Build("POST", "/v2/algo-orders/cancellation", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := algoorder.CancelOrdersResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *AlgoOrderClient) GetOpenOrders(request *model.GetRequest) (*algoorder.GetOpenOrdersResponse, error) {
	url := p.privateUrlBuilder.Build("GET", "/v2/algo-orders/opening", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := algoorder.GetOpenOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *AlgoOrderClient) GetHistoryOrders(request *model.GetRequest) (*algoorder.GetHistoryOrdersResponse, error) {
	url := p.privateUrlBuilder.Build("GET", "/v2/algo-orders/history", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := algoorder.GetHistoryOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *AlgoOrderClient) GetSpecificOrder(request *model.GetRequest) (*algoorder.GetSpecificOrderResponse, error) {
	url := p.privateUrlBuilder.Build("GET", "/v2/algo-orders/specific", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := algoorder.GetSpecificOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

// 自动撤销订单
func (p *AlgoOrderClient) CancelAllAfter(request algoorder.CancelAllAfterRequest) (*algoorder.CancelAllAfterResponse, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/algo-orders/cancel-all-after", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := algoorder.CancelAllAfterResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 {
		return &result, nil
	}

	return nil, errors.New(postResp)
}
