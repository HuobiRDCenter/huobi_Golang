package account

type GetAccountBalanceResponse struct {
	Status string          `json:"status"`
	Data   *AccountBalance `json:"data"`
}

type AccountBalance struct {
	Id    int               `json:"id"`
	Type  string            `json:"type"`
	State string            `json:"state"`
	List  []CurrencyBalance `json:"list"`
}

type CurrencyBalance struct {
	Currency  string `json:"currency"`
	Type      string `json:"type"`
	Balance   string `json:"balance"`
	Debt      string `json:"debt,omitempty"`
	Available string `json:"available,omitempty"`
	SeqNum    string `json:"seq-num,omitempty"`
}
