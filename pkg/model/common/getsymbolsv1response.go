package common

import "github.com/shopspring/decimal"

type GetSymbolsV1Response struct {
	Status  string      `json:"status"`
	Data    []SymbolsV1 `json:"data"`
	Ts      string      `json:"ts"`
	Full    int         `json:"full"`
	ErrCode string      `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
}

type SymbolsV1 struct {
	Symbol  string          `json:"symbol"`
	Sn      string          `json:"sn"`
	Bc      string          `json:"bc"`
	Qc      string          `json:"qc"`
	State   string          `json:"state"`
	Ve      bool            `json:"ve"`
	We      bool            `json:"we"`
	Dl      bool            `json:"dl"`
	Cd      bool            `json:"cd"`
	Te      bool            `json:"te"`
	Ce      bool            `json:"ce"`
	Tet     int64           `json:"tet"`
	Toa     int64           `json:"toa"`
	Tca     int64           `json:"tca"`
	Voa     int64           `json:"voa"`
	Vca     int64           `json:"vca"`
	Sp      string          `json:"sp"`
	Tm      string          `json:"tm"`
	W       int             `json:"w"`
	Ttp     decimal.Decimal `json:"ttp"`
	Tap     decimal.Decimal `json:"tap"`
	Tpp     decimal.Decimal `json:"tpp"`
	Fp      decimal.Decimal `json:"fp"`
	Tags    string          `json:"tags"`
	D       string          `json:"d"`
	Bcdn    string          `json:"bcdn"`
	Qcdn    string          `json:"qcdn"`
	Elr     string          `json:"elr"`
	Castate string          `json:"castate"`
	Ca1oa   int64           `json:"ca1oa"`
	Ca1ca   int64           `json:"ca1ca"`
	Ca2oa   int64           `json:"ca2oa"`
	Ca2ca   int64           `json:"ca2ca"`
}
