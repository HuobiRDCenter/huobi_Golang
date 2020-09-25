package subuser

type SetSubUserTradableMarketRequest struct {
	SubUids     string  `json:"subUids"`
	AccountType string `json:"accountType"`
	Activation  string `json:"activation"`
}
