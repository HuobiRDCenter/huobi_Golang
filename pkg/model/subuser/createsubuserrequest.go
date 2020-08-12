package subuser

type CreateSubUserRequest struct {
	UserList   []Users `json:"userList"`
}


type Users struct {
	UserName  string `json:"userName"`
	Note  string `json:"note"`
}