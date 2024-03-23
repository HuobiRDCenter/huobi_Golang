package common

type GetChainsResponse struct {
	Status  string     `json:"status"`
	Data    []ChainsV1 `json:"data"`
	Ts      string     `json:"ts"`
	Full    int        `json:"full"`
	ErrCode string     `json:"err_code"`
	ErrMsg  string     `json:"err_msg"`
}

type ChainsV1 struct {
	Adt                          bool   `json:"adt"`
	Ac                           string `json:"ac"`
	Ao                           bool   `json:"ao"`
	Awt                          bool   `json:"awt"`
	Chain                        string `json:"chain"`
	Ct                           string `json:"ct"`
	Code                         string `json:"code"`
	Currency                     string `json:"currency"`
	DepositDesc                  string `json:"deposit-desc"`
	De                           bool   `json:"de"`
	Dma                          string `json:"dma"`
	DepositTipsDesc              string `json:"deposit-tips-desc"`
	Dn                           string `json:"dn"`
	Fc                           int    `json:"fc"`
	Ft                           string `json:"ft"`
	Default                      int    `json:"default"`
	ReplaceChainInfoDesc         string `json:"replace-chain-info-desc"`
	ReplaceChainNotificationDesc string `json:"replace-chain-notification-desc"`
	ReplaceChainPopupDesc        string `json:"replace-chain-popup-desc"`
	Ca                           string `json:"ca"`
	Cct                          int    `json:"cct"`
	Sc                           int    `json:"sc"`
	Sda                          string `json:"sda"`
	SuspendDepositDesc           string `json:"suspend-deposit-desc"`
	Swa                          string `json:"swa"`
	SuspendWithdrawDesc          string `json:"suspend-withdraw-desc"`
	V                            bool   `json:"v"`
	WithdrawDesc                 string `json:"withdraw-desc"`
	We                           bool   `json:"we"`
	Wma                          string `json:"wma"`
	Wp                           int    `json:"wp"`
	Fn                           string `json:"fn"`
	WithdrawTipsDesc             string `json:"withdraw-tips-desc"`
	SuspendVisibleDesc           string `json:"suspend-visible-desc"`
}
