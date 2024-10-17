package common

type GetMarketStatusResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    MarketStatus `json:"data"`
}

type MarketStatus struct {
	MarketStatus    int    `json:"marketStatus"`
	HaltStartTime   int64  `json:"haltStartTime"`
	HaltEndTime     int64  `json:"haltEndTime"`
	HaltReason      int    `json:"haltReason"`
	AffectedSymbols string `json:"affectedSymbols"`
}