package account

type FeeSwitchResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    *struct {
	} `json:"data"`
}
