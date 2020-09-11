package order

type RequestOrdersV1Response struct {
	Op        string  `json:"op"`
	Timestamp int64   `json:"ts"`
	Topic     string  `json:"topic"`
	ClientId  string  `json:"cid"`
	ErrorCode int     `json:"err-code"`
	Data      []Order `json:"data"`
}
