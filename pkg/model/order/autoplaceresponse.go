package order

type AutoPlaceResponse struct {
	Code    int       `json:"code"`
	Success bool      `json:"success"`
	Message string    `json:"message,omitempty"`
	Data    AutoPlace `json:"data,omitempty"`
}

type AutoPlace struct {
	OrderID int64 `json:"order-id"`
}
