package margin

type IsolatedMarginLoanOrdersResponse struct {
	Status string                    `json:"status"`
	Data   []IsolatedMarginLoanOrder `json:"data"`
}
type IsolatedMarginLoanOrder struct {
	Id               int64  `json:"id"`
	UserId           int64  `json:"user-id"`
	AccountId        int64  `json:"account-id"`
	Symbol           string `json:"symbol"`
	Currency         string `json:"currency"`
	LoanAmount       string `json:"loan-amount"`
	LoanBalance      string `json:"loan-balance"`
	InterestRate     string `json:"interest-rate"`
	InterestAmount   string `json:"interest-amount"`
	InterestBalance  string `json:"interest-balance"`
	CreatedAt        int64  `json:"created-at"`
	AccruedAt        int64  `json:"accrued-at"`
	State            string `json:"state"`
	PaidPoint        string `json:"paid-point"`
	PaidCoin         string `json:"paid-coin"`
	DeductRate       string `json:"deduct-rate"`
	DeductCurrency   string `json:"deduct-currency"`
	DeductAmount     string `json:"deduct-amount"`
	UpdatedAt        string `json:"updated-at"`
	HourInterestRate string `json:"hour-interest-rate"`
	DayInterestRate  string `json:"day-interest-rate"`
}
