package account

type GetPointBalanceResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    *struct {
		AccountId      string `json:"accountId"`
		AccountStatus  string `json:"accountStatus"`
		AccountBalance string `json:"acctBalance"`
		GroupIds       []struct {
			GroupId      int64  `json:"groupId"`
			ExpiryDate   int64  `json:"expiryDate"`
			RemainAmount string `json:"remainAmount"`
		}
	}
}
