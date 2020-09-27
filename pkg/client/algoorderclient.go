package client

import (
	"encoding/json"
	"github.com/newgoo/huobi_golang/internal"
	"github.com/newgoo/huobi_golang/internal/requestbuilder"
	"github.com/newgoo/huobi_golang/pkg/model"
	"github.com/newgoo/huobi_golang/pkg/model/algoorder"
)

// Responsible to operate algo order
type AlgoOrderClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

// Initializer
func (p *AlgoOrderClient) Init(accessKey string, secretKey string, host string) *AlgoOrderClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
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