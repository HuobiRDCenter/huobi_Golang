package stablecoin

type GetExchangeRateResponse struct {
	Status string `json:"status"`
	Data   *struct {
		Currency       string `json:"currency"`
		Amount         string `json:"amount"`
		Type           string `json:"type"`
		ExchangeAmount string `json:"exchangeAmount"`
		ExchangeFee    string `json:"exchangeFee"`
		QuoteId        string `json:"quoteId"`
		Expiration     string `json:"expiration"`
	}
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
}
