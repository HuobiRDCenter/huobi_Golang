package algoorder

type PlaceOrderRequest struct {
	AccountId     int    `json:"accountId"`
	Symbol        string `json:"symbol"`
	OrderPrice    string `json:"orderPrice"`
	OrderSide     string `json:"orderSide"`
	OrderSize     string `json:"orderSize"`
	OrderValue    string `json:"orderValue"`
	TimeInForce   string `json:"timeInForce"`
	OrderType     string `json:"orderType"`
	ClientOrderId string `json:"clientOrderId"`
	StopPrice     string `json:"stopPrice"`
	TrailingRate  string `json:"trailingRate"`
}