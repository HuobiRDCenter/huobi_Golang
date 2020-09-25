package account

type TransferPointRequest struct {
	FromUid string `json:"fromUid"`
	ToUid   string `json:"toUid"`
	GroupId int64  `json:"groupId"`
	Amount  string `json:"amount"`
}