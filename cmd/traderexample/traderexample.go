package traderexample

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/pkg/client/marketwebsocketclient"
	"github.com/huobirdcenter/huobi_golang/pkg/response/market"
)

func RunAllExamples() {
	subMultipleBBO()
}


func subMultipleBBO() {
	client := new(marketwebsocketclient.BestBidOfferWebSocketClient).Init(config.Host)

	client.SetHandler(
		func() {
			go client.Subscribe("btcusdt", "")
			go client.Subscribe("etcusdt", "")
			go client.Subscribe("bchusdt", "")
			go client.Subscribe("bsvusdt", "")
			go client.Subscribe("dashusdt", "")
			go client.Subscribe("zecusdt", "")
		},
		func(resp interface{}) {
			bboResponse, ok := resp.(market.SubscribeBestBidOfferResponse)
			if ok {
				if bboResponse.Tick != nil {
					t := bboResponse.Tick
					fmt.Printf("Received update, symbol: %s, ask: [%v, %v], bid: [%v, %v]\n", t.Symbol, t.Ask, t.AskSize, t.Bid, t.BidSize)
				}
			}

		})

	err := client.Connect(true)
	if err != nil {
		fmt.Printf("Connect error: %s\n", err)
		return
	}

	fmt.Println("Press ENTER to stop...")
	fmt.Scanln()

	client.Close()
	fmt.Println("Connection closed")
}