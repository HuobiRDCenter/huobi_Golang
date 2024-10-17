package subuser

type CreateSubUserRequest struct {
	UserList       []Users `json:"userList"`
	SubAccountType string  `json:"subAccountType"`
}

type Users struct {
	UserName string `json:"userName"`
	Note     string `json:"note"`
}
