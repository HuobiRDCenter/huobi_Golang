package order

import "github.com/huobirdcenter/huobi_golang/pkg/model/base"

type SubscribeTradeClearResponse struct {
	base.WebSocketV2ResponseBase
	Data *struct {
		Event           string `json:"event"`
		Symbol          string `json:"symbol"`
		OrderId         int64  `json:"orderId"`
		TradePrice      string `json:"tradePrice"`
		TradeVolume     string `json:"tradeVolume"`
		OrderSide       string `json:"orderSide"`
		OrderType       string `json:"orderType"`
		Aggressor       bool   `json:"aggressor"`
		TradeId         int64  `json:"tradeId"`
		TradeTime       int64  `json:"tradeTime"`
		TransactFee     string `json:"transactFee"`
		FeeCurrency     string `json:"feeCurrency"`
		FeeDeduct       string `json:"feeDeduct"`
		FeeDeductType   string `json:"feeDeductType"`
		AccountId       int64  `json:"accountId"`
		Source          string `json:"source"`
		OrderPrice      string `json:"orderPrice"`
		OrderSize       string `json:"orderSize"`
		OrderValue      string `json:"orderValue"`
		ClientOrderId   string `json:"clientOrderId"`
		StopPrice       string `json:"stopPrice"`
		Operator        string `json:"operator"`
		OrderCreateTime int64  `json:"orderCreateTime"`
		OrderStatus     string `json:"orderStatus"`
	} `json:"data"`
}
