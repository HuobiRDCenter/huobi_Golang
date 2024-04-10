package subuser

type GetAccountListResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    AccountList `json:"data,omitempty"`
}

type AccountList struct {
	Uid        int64         `json:"uid"`
	DeductMode string        `json:"deductMode"`
	List       []AccountInfo `json:"list,omitempty"`
}

type AccountInfo struct {
	AccountType   string      `json:"accountType"`
	Activation    string      `json:"activation"`
	Transferrable bool        `json:"transferrable,omitempty"`
	AccountIds    []AccountID `json:"accountIds,omitempty"`
}

type AccountID struct {
	AccountID     int64  `json:"accountId"`
	SubType       string `json:"subType,omitempty"`
	AccountStatus string `json:"accountStatus,omitempty"`
}
