package algoorder

type CancelOrdersRequest struct {
	ClientOrderIds []string `json:"clientOrderIds"`
}
