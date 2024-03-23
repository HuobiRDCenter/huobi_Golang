package algoorder

type CancelAllAfterResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message,omitempty"`
	Data    CancelAllAfter `json:"data"`
}

type CancelAllAfter struct {
	CurrentTime int64 `json:"currentTime"`
	TriggerTime int64 `json:"triggerTime"`
}
