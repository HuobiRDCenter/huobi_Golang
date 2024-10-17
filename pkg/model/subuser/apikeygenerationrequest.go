package subuser

type ApiKeyGenerationRequest struct {
	OtpToken    string `json:"otpToken,omitempty"`
	SubUid      int64  `json:"subUid"`
	Note        string `json:"note"`
	Permission  string `json:"permission"`
	IPAddresses string `json:"ipAddresses,omitempty"`
}
