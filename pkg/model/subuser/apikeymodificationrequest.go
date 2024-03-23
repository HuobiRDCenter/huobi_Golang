package subuser

type ApiKeyModificationRequest struct {
	SubUid      int64  `json:"subUid"`
	AccessKey   string `json:"accessKey"`
	Note        string `json:"note,omitempty"`
	Permission  string `json:"permission,omitempty"`
	IPAddresses string `json:"ipAddresses,omitempty"`
}
