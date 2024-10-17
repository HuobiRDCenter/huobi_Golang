package subuser

type ActiveCreditResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   bool   `json:"data"`
}
