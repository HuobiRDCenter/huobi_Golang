package margin

type CrossMarginGeneralReplyLoanResponse struct {
	Code    int                            `json:"code"`
	Message string                         `json:"message"`
	Data    []CrossMarginGeneraReplaylLoan `json:"data"`
}
type CrossMarginGeneraReplaylLoan struct {
	RepayId   int64 `json:"repayId"`
	RepayTime int64 `json:"repayTime"`
}
