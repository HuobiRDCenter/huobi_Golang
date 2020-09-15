package algoorder

type CancelOrdersResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    CancelledResult `json:"data"`
}

type CancelledResult struct {
	Accepted  []string `json:"accepted"`
	Rejected []string `json:"rejected"`
}