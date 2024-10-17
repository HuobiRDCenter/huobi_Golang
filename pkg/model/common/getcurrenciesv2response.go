package common

type GetCurrenciesV2Response struct {
	Status  string         `json:"status"`
	Data    []CurrenciesV2 `json:"data"`
	Ts      string         `json:"ts"`
	Full    int            `json:"full"`
	ErrCode string         `json:"err_code"`
	ErrMsg  string         `json:"err_msg"`
}

type CurrenciesV2 struct {
	Cc    string `json:"cc"`
	Dn    string `json:"dn"`
	Fn    string `json:"fn"`
	At    int    `json:"at"`
	Wp    int    `json:"wp"`
	Ft    string `json:"ft"`
	Dma   string `json:"dma"`
	Wma   string `json:"wma"`
	Sp2   string `json:"sp2"`
	W2    string `json:"w2"`
	Qc    bool   `json:"qc"`
	State string `json:"state"`
	V     bool   `json:"v"`
	Whe   bool   `json:"whe"`
	Cd    bool   `json:"cd"`
	De    bool   `json:"de"`
	Wed   bool   `json:"wed"`
	Cawt  bool   `json:"cawt"`
	Fc    int    `json:"fc"`
	Sc    int    `json:"sc"`
	Swd   string `json:"swd"`
	Wd    string `json:"wd"`
	Sdd   string `json:"sdd"`
	Dd    string `json:"dd"`
	Svd   string `json:"svd"`
	Tags  string `json:"tags"`
}
