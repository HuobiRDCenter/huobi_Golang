package account

type FeeSwitchRequest struct {
	SwitchType        int    `json:"switchType"`
	DeductionCurrency string `json:"deductionCurrency"`
}
