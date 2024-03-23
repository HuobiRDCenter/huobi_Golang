package orderclientexample

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/order"
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
	autoPlace()
}

func placeOrder() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := order.PlaceOrderRequest{
		AccountId: config.AccountId,
		Type:      "buy-limit",
		Source:    "spot-api",
		Symbol:    "btcusdt",
		Price:     "1.1",
		Amount:    "1",
	}
	resp, err := client.PlaceOrder(&request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			applogger.Info("Place order successfully, order id: %s", resp.Data)
		case "error":
			applogger.Error("Place order error: %s", resp.ErrorMessage)
		}
	}
}

func placeOrders() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := order.PlaceOrderRequest{
		AccountId: config.AccountId,
		Type:      "buy-limit",
		Source:    "spot-api",
		Symbol:    "btcusdt",
		Price:     "1.1",
		Amount:    "1",
	}

	var requests []order.PlaceOrderRequest
	requests = append(requests, request)
	requests = append(requests, request)
	resp, err := client.PlaceOrders(requests)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, r := range resp.Data {
					if r.OrderId != 0 {
						applogger.Info("Place order successfully: order id %d", r.OrderId)
					} else {
						applogger.Info("Place order error: %s", r.ErrorMessage)
					}
				}
			}
		case "error":
			applogger.Error("Place order error: %s", resp.ErrorMessage)
		}
	}
}

func cancelOrderById() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.CancelOrderById("1")
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			applogger.Info("Cancel order successfully, order id: %s", resp.Data)
		case "error":
			applogger.Info("Cancel order error: %s", resp.ErrorMessage)
		}
	}
}

func cancelOrderByClient() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.CancelOrderByClientOrderId("1")
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			applogger.Info("Cancel order successfully, order id: %d", resp.Data)
		case "error":
			applogger.Info("Cancel order error: %s", resp.ErrorMessage)
		}
	}
}

func getOpenOrders() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(model.GetRequest).Init()
	request.AddParam("account-id", config.AccountId)
	request.AddParam("symbol", "btcusdt")
	resp, err := client.GetOpenOrders(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, o := range resp.Data {
					applogger.Info("Open orders, symbol: %s, price: %s, amount: %s", o.Symbol, o.Price, o.Amount)
				}
				applogger.Info("There are total %d open orders", len(resp.Data))
			}
		case "error":
			applogger.Error("Get open order error: %s", resp.ErrorMessage)
		}
	}
}

func cancelOrdersByCriteria() {
	request := order.CancelOrdersByCriteriaRequest{
		AccountId: config.AccountId,
		Symbol:    "btcusdt",
	}

	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.CancelOrdersByCriteria(&request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				d := resp.Data
				applogger.Info("Cancel orders successfully, success count: %d, failed count: %d, next id: %d", d.SuccessCount, d.FailedCount, d.NextId)
			}
		case "error":
			applogger.Error("Cancel orders error: %s", resp.ErrorMessage)
		}
	}
}

func cancelOrdersByIds() {
	request := order.CancelOrdersByIdsRequest{
		OrderIds: []string{"1", "2"},
	}

	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.CancelOrdersByIds(&request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				if resp.Data.Success != nil {
					for _, id := range resp.Data.Success {
						applogger.Info("Cancel orders successfully, id: %s", id)
					}
				}
				if resp.Data.Failed != nil {
					for _, f := range resp.Data.Failed {
						id := f.OrderId
						if id == "" {
							id = f.ClientOrderId
						}
						applogger.Error("Cancel orders failed, id: %s, error: %s", id, f.ErrorMessage)
					}
				}
			}
		case "error":
			applogger.Error("Cancel orders error: %s", resp.ErrorMessage)
		}
	}
}

func getOrderById() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetOrderById("1")
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				o := resp.Data
				applogger.Info("Get order, symbol: %s, price: %s, amount: %s, filled amount: %s, filled cash amount: %s, filled fees: %s",
					o.Symbol, o.Price, o.Amount, o.FilledAmount, o.FilledCashAmount, o.FilledFees)
			}
		case "error":
			applogger.Error("Get order by id error: %s", resp.ErrorMessage)
		}
	}
}

func getOrderByCriteria() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(model.GetRequest).Init()
	request.AddParam("clientOrderId", "cid12345")
	resp, err := client.GetOrderByCriteria(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				o := resp.Data
				applogger.Info("Get order, symbol: %s, price: %s, amount: %s, filled amount: %s, filled cash amount: %s, filled fees: %s",
					o.Symbol, o.Price, o.Amount, o.FilledAmount, o.FilledCashAmount, o.FilledFees)
			}
		case "error":
			applogger.Error("Get order by criteria error: %s", resp.ErrorMessage)
		}
	}
}

func getMatchResultById() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetMatchResultsById("63403286375")
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, r := range resp.Data {
					applogger.Info("Match result, symbol: %s, filled amount: %s, filled fees: %s", r.Symbol, r.FilledAmount, r.FilledFees)
				}
				applogger.Info("There are total %d match results", len(resp.Data))
			}
		case "error":
			applogger.Error("Get match results error: %s", resp.ErrorMessage)
		}
	}
}

func getHistoryOrders() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(model.GetRequest).Init()
	request.AddParam("symbol", "btcusdt")
	request.AddParam("states", "canceled")
	resp, err := client.GetHistoryOrders(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, o := range resp.Data {
					applogger.Info("Order history, symbol: %s, price: %s, amount: %s, state: %s", o.Symbol, o.Price, o.Amount, o.State)
				}
				applogger.Info("There are total %d orders", len(resp.Data))
			}
		case "error":
			applogger.Error("Get history order error: %s", resp.ErrorMessage)
		}
	}
}

func getLast48hOrders() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(model.GetRequest).Init()
	request.AddParam("symbol", "btcusdt")
	resp, err := client.GetLast48hOrders(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, o := range resp.Data {
					applogger.Info("Order history, symbol: %s, price: %s, amount: %s, state: %s", o.Symbol, o.Price, o.Amount, o.State)
				}
				applogger.Info("There are total %d orders", len(resp.Data))
			}
		case "error":
			applogger.Error("Get history order error: %s", resp.ErrorMessage)
		}
	}
}

func getMatchResultByCriteria() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(model.GetRequest).Init()
	request.AddParam("symbol", "btcusdt")
	resp, err := client.GetMatchResultsByCriteria(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			if resp.Data != nil {
				for _, r := range resp.Data {
					applogger.Info("Match result, symbol: %s, filled amount: %s, filled fees: %s", r.Symbol, r.FilledAmount, r.FilledFees)
				}
				applogger.Info("There are total %d match results", len(resp.Data))
			}
		case "error":
			applogger.Error("Get match results error: %s", resp.ErrorMessage)
		}
	}
}

func getTransactFeeRate() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := new(model.GetRequest).Init()
	request.AddParam("symbols", "btcusdt,eosht")
	resp, err := client.GetTransactFeeRate(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Code {
		case 200:
			if resp.Data != nil {
				for _, f := range resp.Data {
					applogger.Info("Fee rate , symbol: %s, maker-taker fee: %s-%s, actual maker-taker fee: %s-%s",
						f.Symbol, f.MakerFeeRate, f.TakerFeeRate, f.ActualMakerRate, f.ActualTakerRate)
				}
				applogger.Info("There are total %d fee rate result", len(resp.Data))
			}
		default:
			applogger.Error("Get fee error: %s", resp.Message)
		}
	}
}

func autoPlace() {
	client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := order.AutoPlaceRequest{Symbol: "btcusdt", AccountID: "31253990", Amount: "0", MarketAmount: "10", Type: "buy-market", TradePurpose: "2", Source: "super-margin-web"}
	resp, err := client.AutoPlace(request)
	if err != nil {
		applogger.Error("autoPlace error: %s", err)
	} else {
		applogger.Info("autoPlace, %v", resp.Data)
	}
}
