package account

type GetAccountAssetValuationResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *struct {
		Balance   string `json:"balance"`
		Timestamp int64  `json:"timestamp"`
	} `json:"data"`
}
