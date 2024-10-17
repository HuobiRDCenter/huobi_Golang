package subuser

type DeductModeResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    []DeductMode `json:"data"`
}

type DeductMode struct {
	SubUid     string `json:"subUid"`
	DeductMode string `json:"deductMode"`
	ErrCode    int64  `json:"errCode,omitempty"`
	ErrMessage string `json:"errMessage,omitempty"`
}
