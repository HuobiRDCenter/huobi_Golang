package market

import (
	"../base"
)

type SubscribeDepthResponse struct {
	base.WebSocketResponseBase
	Data *Depth
	Tick *Depth
}