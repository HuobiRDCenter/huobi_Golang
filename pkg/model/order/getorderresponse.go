package order

type GetOrderResponse struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	Data         *struct {
		AccountId        int    `json:"account-id"`
		Amount           string `json:"amount"`
		Id               int64  `json:"id"`
		ClientOrderId    string `json:"client-order-id"`
		Symbol           string `json:"symbol"`
		Price            string `json:"price"`
		CreatedAt        int64  `json:"created-at"`
		Type             string `json:"type"`
		FilledAmount     string `json:"field-amount"`
		FilledCashAmount string `json:"field-cash-amount"`
		FilledFees       string `json:"field-fees"`
		Source           string `json:"source"`
		State            string `json:"state"`
		FinishedAt       string `json:"finished-at"`
		StopPrice        string `json:"stop-price"`
		Operator         string `json:"operator"`
	}
}
