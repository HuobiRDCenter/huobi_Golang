package margin

type GetMarginLimitResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message,omitempty"`
	Status  string        `json:"status"`
	Data    []MarginLimit `json:"data,omitempty"`
}

type MarginLimit struct {
	Currency    string `json:"currency"`
	MaxHoldings string `json:"maxHoldings"`
}
