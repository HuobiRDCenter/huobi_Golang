package subuser

type SetSubUserTransferabilityRequest struct {
	SubUids int64 `json:"subUids"`
	AccountType string `json:"accountType"`
	Transferrable bool `json:"transferrable"`
}
