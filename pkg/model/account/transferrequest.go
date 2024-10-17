package account

type TransferRequest struct {
	From          string  `json:"from"`
	To            string  `json:"to"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	MarginAccount string  `json:"margin-account"`
}
