package order

import "github.com/huobirdcenter/huobi_golang/pkg/model/base"

type SubscribeOrderV2Response struct {
	base.WebSocketV2ResponseBase
	Data *struct {
		EventType       string `json:"eventType"`
		Symbol          string `json:"symbol"`
		AccountId       int64  `json:"accountId"`
		OrderId         int64  `json:"orderId"`
		ClientOrderId   string `json:"clientOrderId"`
		OrderSide       string `json:"orderSide"`
		OrderPrice      string `json:"orderPrice"`
		OrderSize       string `json:"orderSize"`
		OrderValue      string `json:"orderValue"`
		Type            string `json:"type"`
		OrderStatus     string `json:"orderStatus"`
		OrderCreateTime int64  `json:"orderCreateTime"`
		TradePrice      string `json:"tradePrice"`
		TradeVolume     string `json:"tradeVolume"`
		TradeId         int64  `json:"tradeId"`
		TradeTime       int64  `json:"tradeTime"`
		Aggressor       bool   `json:"aggressor"`
		RemainAmt       string `json:"remainAmt"`
		LastActTime     int64  `json:"lastActTime"`
		ErrorCode       int    `json:"errCode"`
		ErrorMessage    string `json:"errMessage"`
	}
}