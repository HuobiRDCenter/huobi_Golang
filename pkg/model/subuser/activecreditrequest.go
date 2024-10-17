package subuser

type ActiveCreditRequest struct {
	TransactionId int64   `json:"transactionId"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	AccountId     int64   `json:"accountId"`
	UserId        int64   `json:"userId"`
}
