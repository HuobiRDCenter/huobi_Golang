package margin

type CrossMarginGeneralReplayLoanRecordsOptionalRequest struct {
	RepayId   string   `json:"repayId"`
	AccountId string   `json:"accountId"`
	Currency  string   `json:"currency"`
	StartDate int64    `json:"startDate"`
	EndDate   int64    `json:"endDate"`
	Sort      string   `json:"sort"`
	Limit     int      `json:"limit"`
	FromId    int64    `json:"fromId"`
}
