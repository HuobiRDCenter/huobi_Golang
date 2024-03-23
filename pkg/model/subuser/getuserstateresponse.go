package subuser

type GetUserStateResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message,omitempty"`
	Data    UserState `json:"data"`
}

type UserState struct {
	Uid       int64  `json:"uid"`
	UserState string `json:"userState"`
}
