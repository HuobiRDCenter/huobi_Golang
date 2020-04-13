package marketclientexample

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/getrequest"
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

	optionalRequest := getrequest.GetCandlestickOptionalRequest{Period: getrequest.MIN1, Size: 10}
	resp, err := client.GetCandlestick("btcusdt", optionalRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, kline := range resp {
			fmt.Println("High: ", kline.High, "Low:", kline.Low)
		}
	}
}

//  Get the latest ticker with some important 24h aggregated market data for btcusdt.
func getLast24hCandlestickAskBid() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetLast24hCandlestickAskBid("btcusdt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bid: ", resp.Bid, "Ask: ", resp.Ask)
	}
}

//  Get the latest tickers for all supported pairs
func getLast24hCandlesticks() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetAllSymbolsLast24hCandlesticksAskBid()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, tick := range resp {
			fmt.Printf("Symbol: %s, High: %v, Low: %v, Ask[%v, %v], Bid[%v, %v]\n",
				tick.Symbol, tick.High, tick.Low, tick.Ask, tick.AskSize, tick.Bid, tick.BidSize)
		}
	}
}

//  Get the current order book of the btcusdt.
func getDepth() {
	optionalRequest := getrequest.GetDepthOptionalRequest{10}
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetDepth("btcusdt", getrequest.STEP0, optionalRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, ask := range resp.Asks {
			fmt.Println("ask: ", ask)
		}
		for _, bid := range resp.Bids {
			fmt.Println("bid: ", bid)
		}

	}
}

//  Get the latest trade with btucsdt price, volume, and direction.
func getLatestTrade() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetLatestTrade("btcusdt")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, trade := range resp.Data {
			fmt.Println("Id: ", trade.Id, "Price:", trade.Price)
		}
	}
}

//  Get the most recent trades with btcusdt price, volume, and direction.
func getHistoricalTrade() {
	client := new(client.MarketClient).Init(config.Host)
	optionalRequest := getrequest.GetHistoricalTradeOptionalRequest{5}
	resp, err := client.GetHistoricalTrade("btcusdt", optionalRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, tradeData := range resp {
			for _, trade := range tradeData.Data {
				fmt.Println("price: ", trade.Price)
			}
		}
	}
}

//  Get the summary of trading in the market for the last 24 hours.
func getLast24hCandlestick() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetLast24hCandlestick("btcusdt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Close: ", resp.Close, "Open: ", resp.Open)
	}
}
