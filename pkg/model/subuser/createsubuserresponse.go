package subuser

type CreateSubUserResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []UserData `json:"data"`
}

type UserData struct {
	UserName     string `json:"userName"`
	Note         string `json:"note"`
	Uid          int64  `json:"uid"`
	ErrorCode    string  `json:"errCode"`
	ErrorMessage string `json:"errMessage"`
}
