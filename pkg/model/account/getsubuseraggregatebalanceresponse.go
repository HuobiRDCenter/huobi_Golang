package account

type GetSubUserAggregateBalanceResponse struct {
	Status string             `json:"status"`
	Data   []AggregateBalance `json:"data"`
}
type AggregateBalance struct {
	Currency string `json:"currency"`
	Type     string `json:"type"`
	Balance  string `json:"balance"`
}
