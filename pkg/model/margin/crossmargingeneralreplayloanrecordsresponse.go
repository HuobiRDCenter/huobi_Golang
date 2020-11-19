package margin

type CrossMarginGeneralReplyLoanRecordsResponse struct {
	Code   int                                  `json:"code"`
	Data   []CrossMarginGeneraReplaylLoanRecord `json:"data"`
}
type CrossMarginGeneraReplaylLoanRecord struct {
	RepayId         int64      `json:"repayId"`
	RepayTime       int64      `json:"repayTime"`
	AccountId       int64      `json:"accountId"`
	Currency        string     `json:"currency"`
	RepaidAmount    string     `json:"repaidAmount"`
	TransactIds     *Transact  `json:"transactIds"`
}
type Transact struct {
	TransactId       int64    `json:"transactId"`
	RepaidPrincipal  string   `json:"repaidPrincipal"`
	RepaidInterest   string   `json:"repaidInterest"`
	PaidHt           string   `json:"paidHt"`
	PaidPoint        string   `json:"paidPoint"`
}