package marketwebsocketclientexample

import (
	"../../config"
	"../../pkg/client/marketwebsocketclient"
	"../../pkg/getrequest"
	"../../pkg/response/market"
	"fmt"
)

func RunAllExamples() {
	reqAndSubscribeCandlestick()
	reqAndSubscribeDepth()
	reqAndSubscribeMarketByPrice()
	subscribeBBO()
	reqAndSubscribeTrade()
	reqAndSubscribeLast24hCandlestick()
}

func reqAndSubscribeCandlestick() {

	client := new(marketwebsocketclient.CandlestickWebSocketClient).Init(config.Host)

	client.SetHandler(
		func() {
			err := client.Request("btcusdt", "1min", 1569361140, 1569366420, "2305")
			if err != nil {
				fmt.Printf("Sent error: %s\n", err)
			} else {
				fmt.Println("Sent request")
			}

			err = client.Subscribe("btcusdt", "1min", "2118")
			if err != nil {
				fmt.Printf("Subscribe error: %s\n", err)
			} else {
				fmt.Println("Sent subscription")
			}
		},
		func(resp interface{}) {
			candlestickResponse, ok := resp.(market.SubscribeCandlestickResponse)
			if ok {
				if &candlestickResponse != nil {
					if candlestickResponse.Tick != nil {
						t := candlestickResponse.Tick
						fmt.Printf("Candlestick update, id: %d, count: %d, vol: %v [%v-%v-%v-%v]\n",
							t.Id, t.Count, t.Vol, t.Open, t.Close, t.Low, t.High)
					}

					if candlestickResponse.Data != nil {
						for _, t := range candlestickResponse.Data {
							fmt.Printf("Candlestick data, id: %d, count: %d, vol: %v [%v-%v-%v-%v]\n",
								t.Id, t.Count, t.Vol, t.Open, t.Count, t.Low, t.High)
						}
						fmt.Printf("There are total %d ticks\n", len(candlestickResponse.Data))
					}
				}
			} else {
				fmt.Printf("Unknown response: %v\n", resp)
			}

		})

	err := client.Connect(true)
	if err != nil {
		fmt.Printf("Client connect error: %s\n", err)
		return
	}

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	err = client.UnSubscribe("btcusdt", "1min", "2118")
	if err != nil {
		fmt.Printf("UnSubscribe error: %s\n", err)
	}

	client.Close()
	fmt.Println("Client closed")
}

func reqAndSubscribeDepth() {

	client := new(marketwebsocketclient.DepthWebSocketClient).Init(config.Host)

	client.SetHandler(
		func() {
			err := client.Request("btcusdt", getrequest.STEP4, "1153")
			if err != nil {
				fmt.Printf("Sent error: %s\n", err)
			} else {
				fmt.Println("Sent request")
			}

			err = client.Subscribe("btcusdt", getrequest.STEP4, "1153")
			if err != nil {
				fmt.Printf("Subscribe error: %s\n", err)
			} else {
				fmt.Println("Sent subscription")
			}
		},
		func(resp interface{}) {
			depthResponse, ok := resp.(market.SubscribeDepthResponse)
			if ok {
				if &depthResponse != nil {
					if depthResponse.Tick != nil {
						if depthResponse.Tick.Asks != nil {
							a := depthResponse.Tick.Asks
							for i := len(a)-1; i >= 0; i-- {
								fmt.Printf("%v - %v\n", a[i][0], a[i][1])
							}
						}
						fmt.Println("---ask-bid-update--")
						if depthResponse.Tick.Bids != nil {
							b := depthResponse.Tick.Bids
							for i:= 0; i < len(b); i++ {
								fmt.Printf("%v - %v\n", b[i][0], b[i][1])
							}
						}
						fmt.Println()
					}

					if depthResponse.Data != nil {
						if depthResponse.Data.Asks != nil {
							a := depthResponse.Data.Asks
							for i := len(a)-1; i >= 0; i-- {
								fmt.Printf("%v - %v\n", a[i][0], a[i][1])
							}
						}
						fmt.Println("---ask-bid-data--")
						if depthResponse.Data.Bids != nil {
							b := depthResponse.Data.Bids
							for i:= 0; i < len(b); i++ {
								fmt.Printf("%v - %v\n", b[i][0], b[i][1])
							}
						}
						fmt.Println()
					}
				}
			} else {
				fmt.Printf("Unknown response: %v\n", resp)
			}

		})

	err := client.Connect(true)
	if err != nil {
		fmt.Printf("Client connect error: %s\n", err)
		return
	}

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	err = client.UnSubscribe("btcusdt", "1min", "2118")
	if err != nil {
		fmt.Printf("UnSubscribe error: %s\n", err)
	}

	client.Close()
	fmt.Println("Client closed")
}

func reqAndSubscribeMarketByPrice() {

	client := new(marketwebsocketclient.MarketByPriceWebSocketClient).Init(config.Host)

	client.SetHandler(
		func() {
			err := client.Request("btcusdt", "1437")
			if err != nil {
				fmt.Printf("Sent error: %s\n", err)
			} else {
				fmt.Println("Sent request")
			}

			err = client.Subscribe("btcusdt", "1437")
			if err != nil {
				fmt.Printf("Subscribe error: %s\n", err)
			} else {
				fmt.Println("Sent subscription")
			}
		},
		func(resp interface{}) {
			depthResponse, ok := resp.(market.SubscribeMarketByPriceResponse)
			if ok {
				if &depthResponse != nil {
					if depthResponse.Tick != nil {
						t := depthResponse.Tick
						fmt.Printf("MBP prevSeqNum: %d, seqNum: %d\n", t.PrevSeqNum, t.SeqNum)
						if t.Asks != nil {
							for i := len(t.Asks)-1; i >= 0; i-- {
								fmt.Printf("%v - %v\n", t.Asks[i][0], t.Asks[i][1])
							}
						}
						fmt.Println("---MBP-update--")
						if t.Bids != nil {
							for i:= 0; i < len(t.Bids); i++ {
								fmt.Printf("%v - %v\n", t.Bids[i][0], t.Bids[i][1])
							}
						}
						fmt.Println()
					}

					if depthResponse.Data != nil {
						d := depthResponse.Data
						fmt.Printf("MBP prevSeqNum: %d, seqNum: %d\n", d.PrevSeqNum, d.SeqNum)
						if d.Asks != nil {
							a := d.Asks
							for i := len(a)-1; i >= 0; i-- {
								fmt.Printf("%v - %v\n", a[i][0], a[i][1])
							}
						}
						fmt.Println("---MBP-data--")
						if depthResponse.Data.Bids != nil {
							b := depthResponse.Data.Bids
							for i:= 0; i < len(b); i++ {
								fmt.Printf("%v - %v\n", b[i][0], b[i][1])
							}
						}
						fmt.Println()
					}
				}
			} else {
				fmt.Printf("Unknown response: %v\n", resp)
			}

		})

	err := client.Connect(true)
	if err != nil {
		fmt.Printf("Client connect error: %s\n", err)
		return
	}

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	err = client.UnSubscribe("btcusdt", "1437")
	if err != nil {
		fmt.Printf("UnSubscribe error: %s\n", err)
	}

	client.Close()
	fmt.Println("Client closed")
}

func subscribeBBO() {

	client := new(marketwebsocketclient.BestBidOfferWebSocketClient).Init(config.Host)

	client.SetHandler(
		func() {
			err := client.Subscribe("btcusdt", "2118")
			if err != nil {
				fmt.Printf("Subscribe error: %s\n", err)
				return
			}
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

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	err = client.UnSubscribe("btcusdt", "2118")
	if err != nil {
		fmt.Printf("UnSubscribe error: %s\n", err)
	}

	client.Close()
	fmt.Println("Connection closed")
}

func reqAndSubscribeTrade() {

	client := new(marketwebsocketclient.TradeWebSocketClient).Init(config.Host)

	client.SetHandler(
		func() {
			err := client.Request("btcusdt", "1608")
			if err != nil {
				fmt.Printf("Sent error: %s\n", err)
			} else {
				fmt.Println("Sent request")
			}

			err = client.Subscribe("btcusdt", "1608")
			if err != nil {
				fmt.Printf("Subscribe error: %s\n", err)
			} else {
				fmt.Println("Sent subscription")
			}
		},
		func(resp interface{}) {
			depthResponse, ok := resp.(market.SubscribeTradeResponse)
			if ok {
				if &depthResponse != nil {
					if depthResponse.Tick != nil && depthResponse.Tick.Data != nil {
						for _, t := range depthResponse.Tick.Data {
							fmt.Printf("Trade update, id: %d, price: %v, amount: %v\n", t.TradeId, t.Price, t.Amount)
						}
					}

					if depthResponse.Data != nil {
						for _, t := range depthResponse.Data {
							fmt.Printf("Trade data, id: %d, price: %v, amount: %v\n", t.TradeId, t.Price, t.Amount)
						}
					}
				}
			} else {
				fmt.Printf("Unknown response: %v\n", resp)
			}

		})

	err := client.Connect(true)
	if err != nil {
		fmt.Printf("Client connect error: %s\n", err)
		return
	}

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	err = client.UnSubscribe("btcusdt", "1608")
	if err != nil {
		fmt.Printf("UnSubscribe error: %s\n", err)
	}

	client.Close()
	fmt.Println("Client closed")
}

func reqAndSubscribeLast24hCandlestick() {
	// Initialize a new instance
	client := new(marketwebsocketclient.Last24hCandlestickWebSocketClient).Init(config.Host)

	// Set the callback handlers
	client.SetHandler(
		// Connected handler
		func() {
			err := client.Request("btcusdt", "1608")
			if err != nil {
				fmt.Printf("Sent error: %s\n", err)
			} else {
				fmt.Println("Sent request")
			}

			err = client.Subscribe("btcusdt", "1608")
			if err != nil {
				fmt.Printf("Subscribe error: %s\n", err)
			} else {
				fmt.Println("Sent subscription")
			}
		},
		// Response handler
		func(resp interface{}) {
			candlestickResponse, ok := resp.(market.SubscribeLast24hCandlestickResponse)
			if ok {
				if &candlestickResponse != nil {
					if candlestickResponse.Tick != nil {
						t := candlestickResponse.Tick
						fmt.Printf("Candlestick update, id: %d, count: %v, volume: %v [%v-%v-%v-%v]\n",
							t.Id, t.Count, t.Vol, t.Open, t.Close, t.Low, t.High)
					}

					if candlestickResponse.Data != nil {
						t := candlestickResponse.Data
							fmt.Printf("Candlestick data, id: %d, count: %v, volume: %v [%v-%v-%v-%v]\n",
								t.Id, t.Count, t.Vol, t.Open, t.Close, t.Low, t.High)
						}
					}
				} else {
					fmt.Printf("Unknown response: %v\n", resp)
				}

		})

	// Connect to the server and wait for the handler to handle the response
	err := client.Connect(true)
	if err != nil {
		fmt.Printf("Client connect error: %s\n", err)
		return
	}

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	err = client.UnSubscribe("btcusdt", "1608")
	if err != nil {
		fmt.Printf("UnSubscribe error: %s\n", err)
	}

	client.Close()
	fmt.Println("Client closed")
}
