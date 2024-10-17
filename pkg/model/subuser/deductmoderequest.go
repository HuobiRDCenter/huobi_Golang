package subuser

type DeductModeRequest struct {
	SubUids    int64  `json:"subUids"`
	DeductMode string `json:"deductMode"`
}
