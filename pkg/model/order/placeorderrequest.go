package order

type PlaceOrderRequest struct {
	AccountId        string `json:"account-id"`
	Symbol           string `json:"symbol"`
	Type             string `json:"type"`
	Amount           string `json:"amount"`
	Price            string `json:"price,omitempty"`
	Source           string `json:"source,omitempty"`
	ClientOrderId    string `json:"client-order-id,omitempty"`
	SelfMatchPrevent string `json:"self-match-prevent"`
	StopPrice        string `json:"stop-price,omitempty"`
	Operator         string `json:"operator,omitempty"`
}
