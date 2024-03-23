package subuser

type GetUserListResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []UserList `json:"data"`
	NextId  int64      `json:"nextId,omitempty"`
}

type UserList struct {
	Uid         int64  `json:"uid"`
	UserState   string `json:"userState"`
	SubUserName string `json:"subUserName"`
	Note        string `json:"note"`
}
