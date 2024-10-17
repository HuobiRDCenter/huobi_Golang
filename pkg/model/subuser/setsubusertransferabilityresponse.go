package subuser

type SetSubUserTransferabilityResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    []Transferability `json:"data"`
}

type Transferability struct {
	SubUid        int64 `json:"subUid"`
	AccountType   string `json:"accountType"`
	Transferrable bool   `json:"transferrable"`
	ErrCode       int    `json:"errCode"`
	ErrMessage    string `json:"errMessage"`
}