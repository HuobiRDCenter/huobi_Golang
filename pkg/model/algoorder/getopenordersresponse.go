package algoorder

type GetOpenOrdersResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		AccountId     int    `json:"accountId"`
		Source        string `json:"source"`
		ClientOrderId string `json:"clientOrderId"`
		Symbol        string `json:"symbol"`
		OrderPrice    string `json:"orderPrice"`
		OrderSize     string `json:"orderSize"`
		OrderValue    string `json:"orderValue"`
		OrderSide     string `json:"orderSide"`
		TimeInForce   string `json:"timeInForce"`
		OrderType     string `json:"orderType"`
		StopPrice     string `json:"stopPrice"`
		TrailingRate  string `json:"trailingRate"`
		OrderOrigTime int64  `json:"orderOrigTime"`
		LastActTime   int64  `json:"lastActTime"`
		OrderStatus   string `json:"orderStatus"`
	} `json:"data"`
	NextId int64 `json:"nextId"`
}
