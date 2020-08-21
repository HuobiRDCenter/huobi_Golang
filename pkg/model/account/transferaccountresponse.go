package account

type TransferAccountResponse struct {
	Status string              `json:"status"`
	Data   TransferAccountData `json:"data"`
}

type TransferAccountData struct {
	TransactId   int64 `json:"transact-id"`
	TransactTime int64 `json:"transact-time"`
}