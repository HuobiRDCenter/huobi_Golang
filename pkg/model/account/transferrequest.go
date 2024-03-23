package account

import "github.com/shopspring/decimal"

type TransferRequest struct {
	From          string          `json:"from"`
	To            string          `json:"to"`
	Currency      string          `json:"currency"`
	Amount        decimal.Decimal `json:"amount"`
	MarginAccount string          `json:"margin-account"`
}
