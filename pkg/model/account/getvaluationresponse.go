package account

type GetValuationResponse struct {
	Code int       `json:"code"`
	Data Valuation `json:"data"`
}

type Valuation struct {
	TotalBalance             string                 `json:"totalBalance"`
	TodayProfit              string                 `json:"todayProfit"`
	TodayProfitRate          string                 `json:"todayProfitRate"`
	ProfitAccountBalanceList []ProfitAccountBalance `json:"profitAccountBalanceList"`
	Updated                  Updated                `json:"updated"`
	Success                  bool                   `json:"success"`
}

type ProfitAccountBalance struct {
	DistributionType string  `json:"distributionType"`
	Balance          float64 `json:"balance"`
	Success          bool    `json:"success"`
	AccountBalance   string  `json:"accountBalance"`
}

type Updated struct {
	Success bool  `json:"success"`
	Time    int64 `json:"time"`
}
