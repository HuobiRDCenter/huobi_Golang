package wallet

import "github.com/shopspring/decimal"

type GetWithdrawResponse struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	Address           string          `json:"address"`
	ClientOrderID     string          `json:"client-order-id"`
	AddressTag        string          `json:"address-tag"`
	Amount            decimal.Decimal `json:"amount"`
	BlockchainConfirm int             `json:"blockchain-confirm"`
	Chain             string          `json:"chain"`
	CreatedAt         int64           `json:"created-at"`
	Currency          string          `json:"currency"`
	ErrorCode         string          `json:"error-code"`
	ErrorMsg          string          `json:"error-msg"`
	Fee               decimal.Decimal `json:"fee"`
	FromAddrTag       string          `json:"from-addr-tag"`
	FromAddress       string          `json:"from-address"`
	ID                int64           `json:"id"`
	RequestID         string          `json:"request-id"`
	State             string          `json:"state"`
	TxHash            string          `json:"tx-hash"`
	Type              string          `json:"type"`
	UpdatedAt         int64           `json:"updated-at"`
	UserID            int64           `json:"user-id"`
	WalletConfirm     int             `json:"wallet-confirm"`
	SubType           string          `json:"sub-type"`
}
