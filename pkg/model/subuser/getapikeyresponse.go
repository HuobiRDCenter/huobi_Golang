package subuser

type GetApiKeyResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []ApiKey `json:"data"`
}

type ApiKey struct {
	AccessKey   string `json:"accessKey"`
	Note        string `json:"note"`
	Permission  string `json:"permission"`
	IpAddresses string `json:"ipAddresses,omitempty"`
	ValidDays   int    `json:"validDays"`
	Status      string `json:"status"`
	CreateTime  int64  `json:"createTime"`
	UpdateTime  int64  `json:"updateTime"`
}
