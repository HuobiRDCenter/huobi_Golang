package account

type TransferPointResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success string `json:"success"`
	Data    *struct {
		TransactId   string `json:"transactId"`
		TransactTime int64  `json:"transactTime"`
	}
}
