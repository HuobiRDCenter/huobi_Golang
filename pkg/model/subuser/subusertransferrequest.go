package subuser

import "github.com/shopspring/decimal"

type SubUserTransferRequest struct {
	SubUid        int64           `json:"sub-uid"`
	Currency      string          `json:"currency"`
	Amount        decimal.Decimal `json:"amount"`
	Type          string          `json:"type"`
	ClientOrderId string          `json:"client-order-id"`
}
