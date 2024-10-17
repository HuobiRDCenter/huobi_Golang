package account

type GetUserInfoResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    *struct {
		PointSwitch       int    `json:"pointSwitch"`
		CurrencySwitch    int    `json:"currencySwitch"`
		DeductionCurrency string `json:"deductionCurrency"`
	} `json:"data"`
}
