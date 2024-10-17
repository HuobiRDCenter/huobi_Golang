package wallet

type GetWithdrawAddressResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Currency   string `json:"currency"`
		Chain      string `json:"chain"`
		Address    string `json:"address"`
		AddressTag string `json:"addressTag"`
		Note       string `json:"note"`
	}
	NextId int64 `json:"nextId"`
}
