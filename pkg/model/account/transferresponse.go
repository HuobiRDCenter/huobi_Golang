package account

type TransferResponse struct {
	Success bool   `json:"success"`
	Data    int64  `json:"data"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
