package account

type GetSubUserAccountResponse struct {
	Status string           `json:"status"`
	Data   []SubUserAccount `json:"data"`
}
type SubUserAccount struct {
	Id    int               `json:"id"`
	Type  string            `json:"type"`
	State string            `json:"state"`
	List  []CurrencyBalance `json:"list"`
}
