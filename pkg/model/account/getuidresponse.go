package account

type GetUidResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    int64  `json:"data"`
}
