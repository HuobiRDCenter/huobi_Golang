package subuser

type SetSubUserTradableMarketRequest struct {
	SubUids     int64  `json:"subUids"`
	AccountType string `json:"accountType"`
	Activation  string `json:"activation"`
}
