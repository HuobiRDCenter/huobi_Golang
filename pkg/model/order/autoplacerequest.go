package order

type AutoPlaceRequest struct {
	Symbol       string `json:"symbol"`
	AccountID    string `json:"account-id"`
	Amount       string `json:"amount"`
	MarketAmount string `json:"market-amount,omitempty"`
	BorrowAmount string `json:"borrow-amount,omitempty"`
	Type         string `json:"type"`
	TradePurpose string `json:"trade-purpose"`
	Price        string `json:"price,omitempty"`
	StopPrice    string `json:"stop-price,omitempty"`
	Operator     string `json:"operator,omitempty"`
	Source       string `json:"source"`
}
