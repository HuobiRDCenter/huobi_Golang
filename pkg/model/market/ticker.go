package market

type Ticker struct {
	Amount    float64 `json:"amount"`
	Count     int     `json:"count"`
	Open      float64 `json:"open"`
	Close     float64 `json:"close"`
	Low       float64 `json:"low"`
	High      float64 `json:"high"`
	Vol       float64 `json:"vol"`
	Bid       float64 `json:"bid"`
	BidSize   float64 `json:"bidSize"`
	Ask       float64 `json:"ask"`
	AskSize   float64 `json:"askSize"`
	LastPrice float64 `json:"lastPrice"`
	LastSize  float64 `json:"lastSize"`
}
