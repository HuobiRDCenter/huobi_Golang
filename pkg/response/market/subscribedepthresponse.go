package market

import (
	"github.com/huobirdcenter/huobi_golang/pkg/response/base"
)

type SubscribeDepthResponse struct {
	base.WebSocketResponseBase
	Data *Depth
	Tick *Depth
}
