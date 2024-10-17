package margin

type CrossMarginGeneralReplayLoanOptionalRequest struct {
	AccountId    string  `json:"accountId"`
	Currency     string  `json:"currency"`
	Amount       string  `json:"amount"`
	TransactId   string  `json:"transactId"`
}
