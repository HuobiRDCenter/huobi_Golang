package subuser

type SetSubUserTransferabilityRequest struct {
	SubUids string `json:"subUids"`
	AccountType string `json:"accountType"`
	Transferrable bool `json:"transferrable"`
}
