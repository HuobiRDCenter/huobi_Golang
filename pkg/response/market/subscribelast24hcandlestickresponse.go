package market

import "../base"

type SubscribeLast24hCandlestickResponse struct {
	base.WebSocketResponseBase
	Data *Candlestick
	Tick *Candlestick
}