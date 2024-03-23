package subuser

type ApiKeyModificationResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message,omitempty"`
	Data    ApiKeyModification `json:"data,omitempty"`
}

type ApiKeyModification struct {
	Note        string `json:"note,omitempty"`
	Permission  string `json:"permission,omitempty"`
	IPAddresses string `json:"ipAddresses,omitempty"`
}
