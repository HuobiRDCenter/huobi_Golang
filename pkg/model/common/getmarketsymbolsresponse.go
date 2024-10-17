package common

import "github.com/shopspring/decimal"

type GetMarketSymbolsResponse struct {
	Status  string          `json:"status"`
	Data    []MarketSymbols `json:"data"`
	Ts      string          `json:"ts"`
	Full    int             `json:"full"`
	ErrCode string          `json:"err_code"`
	ErrMsg  string          `json:"err_msg"`
}

type MarketSymbols struct {
	Symbol  string          `json:"symbol"`
	BC      string          `json:"bc"`
	QC      string          `json:"qc"`
	State   string          `json:"state"`
	SP      string          `json:"sp"`
	Tags    string          `json:"tags"`
	LR      decimal.Decimal `json:"lr"`
	SMLR    decimal.Decimal `json:"smlr"`
	PP      int             `json:"pp"`
	AP      int             `json:"ap"`
	VP      int             `json:"vp"`
	MinOA   decimal.Decimal `json:"minoa"`
	MaxOA   decimal.Decimal `json:"maxoa"`
	MinOV   decimal.Decimal `json:"minov"`
	LOMinOA decimal.Decimal `json:"lominoa"`
	LOMaxOA decimal.Decimal `json:"lomaxoa"`
	LOMaxBA decimal.Decimal `json:"lomaxba"`
	LOMaxSA decimal.Decimal `json:"lomaxsa"`
	SMMinOA decimal.Decimal `json:"smminoa"`
	SMMaxOA decimal.Decimal `json:"smmaxoa"`
	BMMaxOV decimal.Decimal `json:"bmmaxov"`
	BLMLT   decimal.Decimal `json:"blmlt"`
	SLMGT   decimal.Decimal `json:"slmgt"`
	MSORMLT decimal.Decimal `json:"msormlt"`
	MBORMLT decimal.Decimal `json:"mbormlt"`
	AT      string          `json:"at"`
	U       string          `json:"u"`
	MFR     decimal.Decimal `json:"mfr"`
	CT      string          `json:"ct"`
	RT      string          `json:"rt"`
	RTHR    decimal.Decimal `json:"rthr"`
	IN      decimal.Decimal `json:"in"`
	MaxOV   decimal.Decimal `json:"maxov"`
	FLR     decimal.Decimal `json:"flr"`
	CAState string          `json:"castate"`
}
