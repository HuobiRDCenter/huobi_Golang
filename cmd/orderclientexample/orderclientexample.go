package orderclientexample

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/getrequest"
	"github.com/huobirdcenter/huobi_golang/pkg/postrequest"
)

func RunAllExamples() {
	placeOrder()
	placeOrders()
	cancelOrderById()
	cancelOrderByClient()
	getOpenOrders()
	cancelOrdersByCriteria()
	cancelOrdersByIds()
	getOrderById()
	getOrderByCriteria()
	getMatchResultById()
	getHistoryOrders()
	getLast48hOrders()
	getMatchResultByCriteria()
	getTransactFeeRate()
}

func placeOrder() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := postrequest.PlaceOrderRequest{
		AccountId: config.AccountId,
		Type:      "buy-limit",
		Source:    "spot-api",
		Symbol:    "btcusdt",
		Price:     "1.1",
		Amount:    "1",
	}
	resp, err := client.PlaceOrder(&request)
	if err != nil {
		fmt.Println(err)
	} else {
		switch resp.Status {
		case "ok":
			fmt.Printf("Place order successfully, order id: %s\n", resp.Data)
		case "error":
			fmt.Printf("Place order error: %s\n", resp.ErrorMessage)
		}
	}
}

func placeOrders() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := postrequest.PlaceOrderRequest{
		AccountId: config.AccountId,
		Type:      "buy-limit",
		Source:    "spot-api",
		Symbol:    "btcusdt",
		Price:     "1.1",
		Amount:    "1",
	}

	var requests []postrequest.PlaceOrderRequest
	requests = append(requests, request)
	requests = append(requests, request)
	resp, err := client.PlaceOrders(requests)
	if err != nil {
		fmt.Println(err)
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, r := range resp.Data {
					if r.OrderId != 0 {
						fmt.Printf("Place order successfully: order id %d\n", r.OrderId)
					} else {
						fmt.Printf("Place order error: %s\n", r.ErrorMessage)
					}
				}
			}
		case "error":
			fmt.Printf("Place order error: %s", resp.ErrorMessage)
		}
	}
}

func cancelOrderById() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.CancelOrderById("1")
	if err != nil {
		fmt.Println(err)
	} else {
		switch resp.Status {
		case "ok":
			fmt.Printf("Cancel order successfully, order id: %s\n", resp.Data)
		case "error":
			fmt.Printf("Cancel order error: %s\n", resp.ErrorMessage)
		}
	}
}

func cancelOrderByClient() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.CancelOrderByClientOrderId("1")
	if err != nil {
		fmt.Println(err)
	} else {
		switch resp.Status {
		case "ok":
			fmt.Printf("Cancel order successfully, order id: %d\n", resp.Data)
		case "error":
			fmt.Printf("Cancel order error: %s\n", resp.ErrorMessage)
		}
	}
}

func getOpenOrders() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(getrequest.GetRequest).Init()
	request.AddParam("account-id", config.AccountId)
	request.AddParam("symbol", "btcusdt")
	resp, err := client.GetOpenOrders(request)
	if err != nil {
		fmt.Println(err)
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, o := range resp.Data {
					fmt.Printf("Open orders, symbol: %s, price: %s, amount: %s\n", o.Symbol, o.Price, o.Amount)
				}
				fmt.Printf("There are total %d open orders\n", len(resp.Data))
			}
		case "error":
			fmt.Printf("Get open order error: %s\n", resp.ErrorMessage)
		}
	}
}

func cancelOrdersByCriteria() {
	request := postrequest.CancelOrdersByCriteriaRequest{
		AccountId: config.AccountId,
		Symbol:    "btcusdt",
	}

	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.CancelOrdersByCriteria(&request)
	if err != nil {
		fmt.Println(err)
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				d := resp.Data
				fmt.Printf("Cancel orders successfully, success count: %d, failed count: %d, next id: %d\n", d.SuccessCount, d.FailedCount, d.NextId)
			}
		case "error":
			fmt.Printf("Cancel orders error: %s\n", resp.ErrorMessage)
		}
	}
}

func cancelOrdersByIds() {
	request := postrequest.CancelOrdersByIdsRequest{
		OrderIds: []string{"1", "2"},
	}

	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.CancelOrdersByIds(&request)
	if err != nil {
		fmt.Println(err)
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				if resp.Data.Success != nil {
					for _, id := range resp.Data.Success {
						fmt.Printf("Cancel orders successfully, id: %s\n", id)
					}
				}
				if resp.Data.Failed != nil {
					for _, f := range resp.Data.Failed {
						id := f.OrderId
						if id == "" {
							id = f.ClientOrderId
						}
						fmt.Printf("Cancel orders failed, id: %s, error: %s\n", id, f.ErrorMessage)
					}
				}
			}
		case "error":
			fmt.Printf("Cancel orders error: %s\n", resp.ErrorMessage)
		}
	}
}

func getOrderById() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetOrderById("1")
	if err != nil {
		fmt.Println(err)
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				o := resp.Data
				fmt.Printf("Get order, symbol: %s, price: %s, amount: %s, filled amount: %s, filled cash amount: %s, filled fees: %s\n",
					o.Symbol, o.Price, o.Amount, o.FilledAmount, o.FilledCashAmount, o.FilledFees)
			}
		case "error":
			fmt.Printf("Get order by id error: %s\n", resp.ErrorMessage)
		}
	}
}

func getOrderByCriteria() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(getrequest.GetRequest).Init()
	request.AddParam("clientOrderId", "cid12345")
	resp, err := client.GetOrderByCriteria(request)
	if err != nil {
		fmt.Println(err)
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				o := resp.Data
				fmt.Printf("Get order, symbol: %s, price: %s, amount: %s, filled amount: %s, filled cash amount: %s, filled fees: %s\n",
					o.Symbol, o.Price, o.Amount, o.FilledAmount, o.FilledCashAmount, o.FilledFees)
			}
		case "error":
			fmt.Printf("Get order by criteria error: %s\n", resp.ErrorMessage)
		}
	}
}

func getMatchResultById() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetMatchResultsById("63403286375")
	if err != nil {
		fmt.Println(err)
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, r := range resp.Data {
					fmt.Printf("Match result, symbol: %s, filled amount: %s, filled fees: %s\n", r.Symbol, r.FilledAmount, r.FilledFees)
				}
				fmt.Printf("There are total %d match results\n", len(resp.Data))
			}
		case "error":
			fmt.Printf("Get match results error: %s\n", resp.ErrorMessage)
		}
	}
}

func getHistoryOrders() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(getrequest.GetRequest).Init()
	request.AddParam("symbol", "btcusdt")
	request.AddParam("states", "canceled")
	resp, err := client.GetHistoryOrders(request)
	if err != nil {
		fmt.Println(err)
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, o := range resp.Data {
					fmt.Printf("Order history, symbol: %s, price: %s, amount: %s, state: %s\n", o.Symbol, o.Price, o.Amount, o.State)
				}
				fmt.Printf("There are total %d orders\n", len(resp.Data))
			}
		case "error":
			fmt.Printf("Get history order error: %s\n", resp.ErrorMessage)
		}
	}
}

func getLast48hOrders() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(getrequest.GetRequest).Init()
	request.AddParam("symbol", "btcusdt")
	resp, err := client.GetLast48hOrders(request)
	if err != nil {
		fmt.Println(err)
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, o := range resp.Data {
					fmt.Printf("Order history, symbol: %s, price: %s, amount: %s, state: %s\n", o.Symbol, o.Price, o.Amount, o.State)
				}
				fmt.Printf("There are total %d orders\n", len(resp.Data))
			}
		case "error":
			fmt.Printf("Get history order error: %s\n", resp.ErrorMessage)
		}
	}
}

func getMatchResultByCriteria() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(getrequest.GetRequest).Init()
	request.AddParam("symbol", "btcusdt")
	resp, err := client.GetMatchResultsByCriteria(request)
	if err != nil {
		fmt.Println(err)
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, r := range resp.Data {
					fmt.Printf("Match result, symbol: %s, filled amount: %s, filled fees: %s\n", r.Symbol, r.FilledAmount, r.FilledFees)
				}
				fmt.Printf("There are total %d match results\n", len(resp.Data))
			}
		case "error":
			fmt.Printf("Get match results error: %s\n", resp.ErrorMessage)
		}
	}
}

func getTransactFeeRate() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(getrequest.GetRequest).Init()
	request.AddParam("symbols", "btcusdt,eosht")
	resp, err := client.GetTransactFeeRate(request)
	if err != nil {
		fmt.Println(err)
	} else {
		switch resp.Code {
		case 200:
			if resp.Data != nil {
				for _, f := range resp.Data {
					fmt.Printf("Fee rate , symbol: %s, maker-taker fee: %s-%s, actual maker-taker fee: %s-%s\n",
						f.Symbol, f.MakerFeeRate, f.TakerFeeRate, f.ActualMakerRate, f.ActualTakerRate)
				}
				fmt.Printf("There are total %d fee rate result\n", len(resp.Data))
			}
		default:
			fmt.Printf("Get fee error: %s\n", resp.Message)
		}
	}
}
