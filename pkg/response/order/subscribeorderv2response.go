package order

type SubscribeOrderV2Response struct {
	Ch string `json:"ch"`
	Data struct {
		Symbol string `json:"symbol"`
		OrderId int64 `json:"orderId"`
		TradePrice string `json:"tradePrice"`
		TradeVolume string `json:"tradeVolume"`
		OrderSide string `json:"orderSide"`
		OrderType string `json:"orderType"`
		Aggressor bool `json:"aggressor"`
		TradeId int64 `json:"tradeId"`
		TradeTime int64 `json:"tradeTime"`
		TransactFee string `json:"transactFee"`
		FeeDeduct string `json:"feeDeduct"`
		FeeDeductType string `json:"feeDeductType"`
	}
}