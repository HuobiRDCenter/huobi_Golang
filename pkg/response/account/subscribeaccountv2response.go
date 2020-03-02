package account

type SubscribeAccountV2Response struct {
	Action string `json:"action"`
	Ch string `json:"ch"`
	Data struct {
		Currency string `json:"currency"`
		AccountId int `json:"accountId"`
		Balance string `json:"balance"`
		ChangeType string `json:"changeType"`
		AccountType string `json:"accountType"`
		ChangeTime int64 `json:"changeTime"`
	}
}