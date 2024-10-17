package market

import (
	"github.com/huobirdcenter/huobi_golang/pkg/model/base"
)

type SubscribeLast24hCandlestickResponse struct {
	base.WebSocketResponseBase
	Data *Candlestick
	Tick *Candlestick
}
