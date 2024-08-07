package wallet

type GetDepositAddressResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    []DepositAddress `json:"data"`
}
type DepositAddress struct {
	UserId     int64  `json:"userId"`
	Currency   string `json:"currency"`
	Address    string `json:"address"`
	AddressTag string `json:"addressTag"`
	Chain      string `json:"chain"`
}
