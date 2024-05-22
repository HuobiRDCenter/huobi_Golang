package account

type GetOverviewInfoResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    *struct {
		Currency string `json:"currency"`
	} `json:"data"`
}
