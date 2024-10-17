package subuser

type ApiKeyGenerationResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message,omitempty"`
	Data    ApiKeyGeneration `json:"data"`
}

type ApiKeyGeneration struct {
	Note        string `json:"note"`
	AccessKey   string `json:"accessKey"`
	SecretKey   string `json:"secretKey"`
	Permission  string `json:"permission"`
	IPAddresses string `json:"ipAddresses,omitempty"`
}
