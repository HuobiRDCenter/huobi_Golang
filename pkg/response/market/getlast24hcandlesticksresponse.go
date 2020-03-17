package market

import "github.com/shopspring/decimal"

type GetLast24hCandlesticksResponse struct {
	Status string              `json:"status"`
	Ts     int64               `json:"ts"`
	Data   []SymbolCandlestick `json:"data"`
}
type SymbolCandlestick struct {
	Amount decimal.Decimal `json:"amount"`
	Open   decimal.Decimal `json:"open"`
	Close  decimal.Decimal `json:"close"`
	High   decimal.Decimal `json:"high"`
	Symbol string          `json:"symbol"`
	Count  int64           `json:"count"`
	Low    decimal.Decimal `json:"low"`
	Vol    decimal.Decimal `json:"vol"`
}
