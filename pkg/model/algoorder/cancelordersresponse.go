package algoorder

type CancelOrdersResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Accepted []string `json:"accepted"`
		Rejected []string `json:"rejected"`
	} `json:"data"`
}
