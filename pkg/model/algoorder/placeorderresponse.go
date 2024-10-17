package algoorder

type PlaceOrderResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *struct {
		ClientOrderId string `json:"clientOrderId"`
	} `json:"data"`
}
