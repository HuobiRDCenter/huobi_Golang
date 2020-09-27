package marketclientexample

import (
	"github.com/newgoo/huobi_golang/config"
	"github.com/newgoo/huobi_golang/logging/applogger"
	"github.com/newgoo/huobi_golang/pkg/client"
	"github.com/newgoo/huobi_golang/pkg/model/market"
)

func RunAllExamples() {
	getCandlestick()
	getLast24hCandlestickAskBid()
	getLast24hCandlesticks()
	getDepth()
	getLatestTrade()
	getHistoricalTrade()
	getLast24hCandlestick()
}

//  Get the candlestick/kline for the btcusdt. The specified data number is 10 .
func getCandlestick() {
	client := new(client.MarketClient).Init(config.Host)

	optionalRequest := market.GetCandlestickOptionalRequest{Period: market.MIN1, Size: 10}
	resp, err := client.GetCandlestick("btcusdt", optionalRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, kline := range resp {
			applogger.Info("High=%v, Low=%v", kline.High, kline.Low)
		}
	}
}

//  Get the latest ticker with some important 24h aggregated market data for btcusdt.
func getLast24hCandlestickAskBid() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetLast24hCandlestickAskBid("btcusdt")
	if err != nil {
		applogger.Error(err.Error())
	} else {
		applogger.Info("Bid=%+v, Ask=%+v", resp.Bid, resp.Ask)
	}
}

//  Get the latest tickers for all supported pairs
func getLast24hCandlesticks() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetAllSymbolsLast24hCandlesticksAskBid()
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, tick := range resp {
			applogger.Info("Symbol: %s, High: %v, Low: %v, Ask[%v, %v], Bid[%v, %v]",
				tick.Symbol, tick.High, tick.Low, tick.Ask, tick.AskSize, tick.Bid, tick.BidSize)
		}
	}
}

//  Get the current order book of the btcusdt.
func getDepth() {
	optionalRequest := market.GetDepthOptionalRequest{10}
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetDepth("btcusdt", market.STEP0, optionalRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, ask := range resp.Asks {
			applogger.Info("ask: %+v", ask)
		}
		for _, bid := range resp.Bids {
			applogger.Info("bid: %+v", bid)
		}

	}
}

//  Get the latest trade with btucsdt price, volume, and direction.
func getLatestTrade() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetLatestTrade("btcusdt")
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, trade := range resp.Data {
			applogger.Info("Id=%v, Price=%v", trade.Id, trade.Price)
		}
	}
}

//  Get the most recent trades with btcusdt price, volume, and direction.
func getHistoricalTrade() {
	client := new(client.MarketClient).Init(config.Host)
	optionalRequest := market.GetHistoricalTradeOptionalRequest{5}
	resp, err := client.GetHistoricalTrade("btcusdt", optionalRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, tradeData := range resp {
			for _, trade := range tradeData.Data {
				applogger.Info("price: %v", trade.Price)
			}
		}
	}
}

//  Get the summary of trading in the market for the last 24 hours.
func getLast24hCandlestick() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetLast24hCandlestick("btcusdt")
	if err != nil {
		applogger.Error(err.Error())
	} else {
		applogger.Info("Close=%v, Open=%v", resp.Close, resp.Open)
	}
}
