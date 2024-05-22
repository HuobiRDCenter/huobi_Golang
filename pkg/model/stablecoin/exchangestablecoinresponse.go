package stablecoin

type ExchangeStableCoinResponse struct {
	Status string `json:"status"`
	Data   *struct {
		TransactId     int64  `json:"transact-id"`
		Currency       string `json:"currency"`
		Amount         string `json:"amount"`
		Type           string `json:"type"`
		ExchangeAmount string `json:"exchange-amount"`
		ExchangeFee    string `json:"exchange-fee"`
		Time           string `json:"time"`
	} `json:"data"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
}
