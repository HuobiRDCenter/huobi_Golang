package common

type GetCurrencysV1Response struct {
	Status  string        `json:"status"`
	Data    []CurrencysV1 `json:"data"`
	Ts      string        `json:"ts"`
	Full    int           `json:"full"`
	ErrCode string        `json:"err_code"`
	ErrMsg  string        `json:"err_msg"`
}

type CurrencysV1 struct {
	Name  string   `json:"name"`
	Dn    string   `json:"dn"`
	Vat   int64    `json:"vat"`
	Det   int64    `json:"det"`
	Wet   int64    `json:"wet"`
	Wp    int      `json:"wp"`
	Ct    string   `json:"ct"`
	Cp    string   `json:"cp"`
	Ss    []string `json:"ss"`
	Oe    int      `json:"oe"`
	Dma   string   `json:"dma"`
	Wma   string   `json:"wma"`
	Sp    string   `json:"sp"`
	W     int64    `json:"w"`
	Qc    bool     `json:"qc"`
	State string   `json:"state"`
	V     bool     `json:"v"`
	Whe   bool     `json:"whe"`
	Cd    bool     `json:"cd"`
	De    bool     `json:"de"`
	We    bool     `json:"we"`
	Cawt  bool     `json:"cawt"`
	Cao   bool     `json:"cao"`
	Fc    int      `json:"fc"`
	Sc    int      `json:"sc"`
	Swd   string   `json:"swd"`
	Wd    string   `json:"wd"`
	Sdd   string   `json:"sdd"`
	Dd    string   `json:"dd"`
	Svd   string   `json:"svd"`
	Tags  string   `json:"tags"`
	Fn    string   `json:"fn"`
	Bc    string   `json:"bc"`
	Iqc   bool     `json:"iqc"`
}
