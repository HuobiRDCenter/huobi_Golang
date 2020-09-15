package algoorder

type PlaceOrderResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    ClientOrderIdData `json:"data"`
}

type ClientOrderIdData struct {
	ClientOrderId string `json:"clientOrderId"`
}
