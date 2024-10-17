package market

import (
	"github.com/huobirdcenter/huobi_golang/pkg/model/base"
	"github.com/shopspring/decimal"
)

type SubscribeTradeResponse struct {
	base.WebSocketResponseBase
	Data []Trade
	Tick *struct {
		Id        int64 `json:"id"`
		Timestamp int64 `json:"ts"`
		Data      []Trade
	}
}

type Trade struct {
	TradeId   int64           `json:"tradeId"`
	Amount    decimal.Decimal `json:"amount"`
	Price     decimal.Decimal `json:"price"`
	Timestamp int64           `json:"ts"`
	Direction string          `json:"direction"`
}
