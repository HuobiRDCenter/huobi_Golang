package account

type SubscribeAccountV1Response struct {
	Op        string `json:"op"`
	Timestamp int64  `json:"ts"`
	Topic     string `json:"topic"`
	ClientId  string `json:"cid"`
	ErrorCode int    `json:"err-code"`
	Data      struct {
		Event string `json:"event"`
		List  []struct {
			AccountId int    `json:"account-id"`
			Currency  string `json:"currency"`
			Type      string `json:"type"`
			Balance   string `json:"balance"`
		}
	}
}
