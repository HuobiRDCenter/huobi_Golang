package subuser

type SetSubUserTradableMarketResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    []TradableMarket `json:"data"`
}

type TradableMarket struct {
	SubUid      string `json:"subUid"`
	AccountType string `json:"accountType"`
	Activation  string `json:"activation"`
	ErrCode     int    `json:"errCode"`
	ErrMessage  string `json:"errMessage"`
}