package common

import "github.com/shopspring/decimal"

type GetSymbolsV2Response struct {
	Status  string     `json:"status"`
	Data    []SymbolV2 `json:"data"`
	Ts      string     `json:"ts"`
	Full    int        `json:"full"`
	ErrCode string     `json:"err_code"`
	ErrMsg  string     `json:"err_msg"`
}

type SymbolV2 struct {
	Si                string          `json:"si"`
	Scr               string          `json:"scr"`
	Sc                string          `json:"sc"`
	Dn                string          `json:"dn"`
	Bc                string          `json:"bc"`
	Bcdn              string          `json:"bcdn"`
	Qc                string          `json:"qc"`
	Qcdn              string          `json:"qcdn"`
	State             string          `json:"state"`
	Whe               bool            `json:"whe"`
	Cd                bool            `json:"cd"`
	Te                bool            `json:"te"`
	Toa               int64           `json:"toa"`
	Sp                string          `json:"sp"`
	W                 int             `json:"w"`
	Ttp               decimal.Decimal `json:"ttp"`
	Tap               decimal.Decimal `json:"tap"`
	Tpp               decimal.Decimal `json:"tpp"`
	Fp                decimal.Decimal `json:"fp"`
	SuspendDesc       string          `json:"suspend_desc"`
	TransferBoardDesc string          `json:"transfer_board_desc"`
	Tags              string          `json:"tags"`
	Lr                decimal.Decimal `json:"lr"`
	Smlr              decimal.Decimal `json:"smlr"`
	Flr               string          `json:"flr"`
	Wr                string          `json:"wr"`
	D                 int             `json:"d"`
	Elr               string          `json:"elr"`
	P                 []P             `json:"p"`
	Castate           string          `json:"castate"`
	Ca1oa             int64           `json:"ca1oa"`
	Ca2oa             int64           `json:"ca2oa"`
}

type P struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}
