# Huobi Golang SDK

This is Huobi Go SDK,  you can install to your Golang project and use this SDK to query all market data, trading and manage your account.

The SDK supports RESTful API invoking, and subscribe the market, account and order update from the Websocket connection.

## Table of Contents

- [Quick Start](#Quick-Start)
- [Usage](#Usage)
  - [Folder structure](#Folder-Structure)
  - [Run examples](#Run-examples)
  - [Client](#Client)
  - [Response](#Response)
  - [Init function](#Init-function)
- [Request examples](#Request-examples)
  - [Common data](#Common-data)
  - [Market data](#Market-data)
  - [Account](#account)
  - [Wallet](#wallet)
  - [Trading](#trading)
  - [Margin Loan](#margin-loan)
- [Subscription examples](#Subscription-examples)
  - [Subscribe trade update](#Subscribe-trade-update)
  - [Subscribe candlestick update](#Subscribe-candlestick-update)
  - [Subscribe order update](#subscribe-order-update)
  - [Subscribe account change](#subscribe-account-change)
- [Unsubscribe](#unsubscribe)

## Quick Start

The SDK is compiled by Go 1.13.7, you can import this SDK in your Golang project:

* Import **/pkg/client** package
* Create one of the clients (under package **/pkg/client**) instance by **Init** method
* Call the method provided by client.

```go
import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
)

// Get the timestamp from Huobi server and print on console
client := new(client.CommonClient).Init(config.Host)
resp, err := client.GetTimestamp()

if err != nil {
  fmt.Println(err)
} else {
  fmt.Println("timestamp:", resp)
}


// Get the list of accounts owned by this API user and print the detail on console
client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
resp, err := client.GetAccountInfo()
if err != nil {
  fmt.Println(err)
} else {
  for _, result := range resp {
    fmt.Printf("account: %+v\n", result)
    }
}
```

## Usage

After above section, this SDK should be already download to your local machine, this section introduce this SDK and how to use it correctly.

### Folder Structure

This is the folder and package structure of SDK source code and the description

- **pkg**: The public package of the SDK
  - **client**: The client struct that are responsible to access data
  - **getrequest**: The get request data model
  - **postrequest**: The post request data model
  - **response**: The response data model
- **internal** The internal package that used internally
  - **gzip**: it provide the gzip decompress functionality that unzip the websocket binary data
  - **model**: The internal data model
  - **requestbuilder**: Responsible to build the request with the signature
- **config**: it stores the common configuration, such as host, access key.
- **cmd**: The main package is defined here, it provides the examples how to use **client** package and **response** package to access API and read response.

As the example indicates, there are two important namespaces: **client** and **response**,  this section will introduce both of them below.

### Run examples

This SDK provides examples that under **/cmd** folder, if you want to run the examples to access private data, you need below additional steps:

1. Create an **API Key** first from Huobi official website
2. Create **key.go** into your **config** folder (package). The purpose of this file is to prevent submitting SecretKey into repository by accident, so this file is already added in the *.gitignore* file. 

3. Assign your secret key to string *SecretKey*:

```go
// key.go file
package config

// replace with your API SecretKey
var SecretKey = "xxxx-xxxx-xxxx-xxxx"
```

If you don't need to access private data, you can ignore the secret key.

Regarding the difference between public data and private data you can find details in [Client](#Client) section below.

### Client

In this SDK, the client is the struct to access the Huobi API. In order to isolate the private data with public data, and isolated different kind of data, the client category is designated to match the API category. 

All the client is listed in below table. Each client is very small and simple, it is only responsible to operate its related data, you can pick up multiple clients to create your own application based on your business.

| Data Category   | Client                            | Privacy | Access Type  |
| --------------- | --------------------------------- | ------- | ------------ |
| Common          | CommonClient                      | Public  | Rest         |
| Market          | MarketClient                      | Public  | Rest         |
|                 | CandlestickWebSocketClient        | Public  | WebSocket v1 |
|                 | DepthWebSocketClient              | Public  | WebSocket v1 |
|                 | MarketByPriceWebSocketClient      | Public  | WebSocket v1 |
|                 | BestBidOfferWebSocketClient       | Public  | WebSocket v1 |
|                 | TradeWebSocketClient              | Public  | WebSocket v1 |
|                 | Last24hCandlestickWebSocketClient | Public  | WebSocket v1 |
| Account         | AccountClient                     | Private | Rest         |
|                 | RequestAccountWebSocketV1Client   | Private | WebSocket v1 |
|                 | SubscribeAccountWebSocketV1Client | Private | WebSocket v1 |
|                 | SubscribeAccountWebSocketV2Client | Private | WebSocket v2 |
| Wallet          | WalletClient                      | Private | Rest         |
| Order           | OrderClient                       | Private | Rest         |
|                 | RequestOrdersWebSocketV1Client    | Private | WebSocket v1 |
|                 | RequestOrderWebSocketV1Client     | Private | WebSocket v1 |
|                 | SubscribeOrderWebSocketV1Client   | Private | WebSocket v1 |
|                 | SubscribeOrderWebSocketV2Client   | Private | WebSocket v2 |
| Isolated Margin | IsolatedMarginClient              | Private | Rest         |
| Cross Margin    | CrossMarginClient                 | Private | Rest         |
| ETF             | ETFClient                         | Private | Rest         |

#### Public vs. Private

There are two types of privacy that is correspondent with privacy of API:

**Public client**: It invokes public API to get public data (Common data and Market data), therefore you can create a new instance without applying an API Key.

```go
// Create a CommonClient instance
client := new(client.CommonClient).Init(config.Host)

// Create a CandlestickWebSocketClient instance
client := new(marketwebsocketclient.CandlestickWebSocketClient).Init(config.Host)
```

**Private client**: It invokes private API to access private data, you need to follow the API document to apply an API Key first, and pass the API Key to the init function

```go
// Create an AccountClient instance with APIKey
client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)

// Create a RequestOrdersWebSocketV1Client instance with API Key
client := new(orderwebsocketclient.RequestOrderWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)
```

The API key is used for authentication. If the authentication cannot pass, the invoking of private interface will fail.

#### Rest vs. WebSocket

**Rest Client**: It invokes Rest API and get once-off response.

**WebSocket Client**: It establishes WebSocket connection with server and data will be pushed from server actively. There are two types of method for WebSocket client:

- Request method: The method name starts with "Request-", it will receive the once-off data after sending the request.
- Subscription: The method name starts with "Subscribe-", it will receive update after sending the subscription.

### Response

In this SDK, the response is the struct that define the data returned from API, which is unmarshal from JSON string. It is the return type from each client method. The Rest client also returns an error, you should check the error before read the response:

```go
// Use 'err' to receive the potential error information
resp, err := client.GetTimestamp()

// Check error first
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println("timestamp:", resp)
}
```

For array struct, you can use for/range to iterate each element

```go
// Check the status of response and print some properties
for _, kline := range resp {
  fmt.Println("High: ", kline.High, "Low:", kline.Low)
}
```

### Init function

Golang is not a pure object oriented programming language, and there is no native constructor. In this SDK, every struct has an ***Init*** function for each struct, you must call Init function first, otherwise the member variables may not be initialized expected.

## Request Examples

### Common data

#### Exchange timestamp

```go
client := new(client.CommonClient).Init(config.Host)
resp, err := client.GetTimestamp()
```

#### Symbol and currencies

```go
client := new(client.CommonClient).Init(config.Host)
resp, err := client.GetSymbols()

resp, err := client.GetCurrencys()
```

### Market data

#### Candlestick/KLine

```go
client := new(client.MarketClient).Init(config.Host)
optionalRequest := getrequest.GetCandlestickOptionalRequest{Period: getrequest.MIN1, Size: 10}
resp, err := client.GetCandlestick("btcusdt", optionalRequest)
```

#### Depth

```go
client := new(client.MarketClient).Init(config.Host)
resp, err := client.GetMarketDetailMerged("btcusdt")
```

#### Latest trade

```go
client := new(client.MarketClient).Init(config.Host)
resp, err := client.GetLatestTrade("btcusdt")
```

#### Best bid/ask

```go
client := new(client.MarketClient).Init(config.Host)
resp, err := client.GetLast24hCandlestickAskBid("btcusdt")
```

#### Historical trade

```go
client := new(client.MarketClient).Init(config.Host)
optionalRequest := getrequest.GetHistoricalTradeOptionalRequest{5}
resp, err := client.GetHistoricalTrade("btcusdt", optionalRequest)
```

#### 24H statistics

```go
client := new(client.MarketClient).Init(config.Host)
resp, err := client.GetLast24hCandlestick("btcusdt")
```

### Account

*Authentication is required.*

```go
client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
resp, err := client.GetAccountInfo()
```

### Wallet

*Authentication is required.*

#### Withdraw

```go
client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
createWithdrawRequest := postrequest.CreateWithdrawRequest{
  Address:  "xxxx",
  Amount:   "1.0",
  Currency: "usdt",
  Fee:      "1.0"}
resp, err := client.CreateWithdraw(createWithdrawRequest)
```

#### Cancel withdraw

```go
client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
resp, err := client.CancelWithdraw(12345)
```

#### Withdraw and deposit history

```go
client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
depositType := "deposit"
queryDepositWithdrawOptionalRequest := getrequest.QueryDepositWithdrawOptionalRequest{Currency: "usdt"}

resp, err := client.QueryDepositWithdraw(depositType, queryDepositWithdrawOptionalRequest)
```

### Trading

*Authentication is required.*

#### Create order

```go
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
```

#### Cancel order

```go
client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
resp, err := client.CancelOrderById("1")
```

#### Cancel open orders

```go
request := postrequest.CancelOrdersByCriteriaRequest{
  AccountId: config.AccountId,
  Symbol:    "btcusdt",
}

client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
resp, err := client.CancelOrdersByCriteria(&request)
```

#### Get order info

```go
client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)	resp, err := client.GetOrderById("1")
```

#### Historical orders

```go
client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
request := new(getrequest.GetRequest).Init()
request.AddParam("symbol", "btcusdt")
request.AddParam("states", "canceled")
resp, err := client.GetHistoryOrders(request)
```

### Margin Loan

*Authentication is required.*

#### Apply loan

```go
client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
request := postrequest.IsolatedMarginOrdersRequest{
  Currency: "eos",
  Amount: "0.001",
  Symbol: "btcusdt",
}
resp, err := client.MarginOrders(request)
```

#### Repay loan

```go
client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
orderId := "12345"
request := postrequest.MarginOrdersRepayRequest{Amount: "1.0"}

resp, err := client.MarginOrdersRepay(orderId, request)
```

#### Loan history

```go
client := new(client.IsolatedMarginClient).Init(config.AccessKey, config.SecretKey, config.Host)
optionalRequest := getrequest.IsolatedMarginLoanOrdersOptionalRequest{
  StartDate: "2020-1-1",
}	

resp, err := client.MarginLoanOrders("btcusdt", optionalRequest)
```

## Subscription Examples

### Subscribe trade update

*Authentication is required.*

```go
// Initialize a new instance
client := new(orderwebsocketclient.SubscribeOrderWebSocketV2Client).Init(config.AccessKey, config.SecretKey, config.Host)

// Set the callback handlers
client.SetHandler(
  // Connected handler
  func(resp *model.WebSocketV2AuthenticationResponse) {
    if resp.IsAuth() {
      // Subscribe if authentication passed
      err := client.Subscribe("1", "1149")
      if err != nil {
        fmt.Printf("Subscribe error: %s\n", err)
      } else {
        fmt.Println("Sent subscription")
      }
    } else {
      fmt.Printf("Authentication error: %d\n", resp.Code)
    }
  },
  // Response handler
  func(resp interface{}) {
    subResponse, ok := resp.(order.SubscribeOrderV2Response)
    if ok {
      if &subResponse.Data != nil {
        o := subResponse.Data
        fmt.Printf("Order update, symbol: %s, order id: %d, price: %s, volume: %s",
                   o.Symbol, o.OrderId, o.TradePrice, o.TradeVolume)
      }
    } else {
      fmt.Printf("Received unknown response: %v\n", resp)
    }
})

// Connect to the server and wait for the handler to handle the response
err := client.Connect(true)
if err != nil {
  fmt.Printf("Client Connect error: %s\n", err)
  return
}
```

### Subscribe candlestick update

```go
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
```

### Subscribe order update

*Authentication is required.*

```go
// Initialize a new instance
client := new(orderwebsocketclient.SubscribeOrderWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)

// Set the callback handlers
client.SetHandler(
  // Connected handler
  func(resp *model.WebSocketV1AuthenticationResponse) {
    if resp.ErrorCode == 0 {
      err := client.Subscribe("btcusdt", "1601")
        if err != nil {
          fmt.Printf("Subscribe error: %s\n", err)
        } else {
          fmt.Println("Sent subscription")
        }
    } else {
      fmt.Printf("Authentication error: %d\n", resp.ErrorCode)
    }
  },
  // Response handler
  func(resp interface{}) {
    subResponse, ok := resp.(order.SubscribeOrderV1Response)
      if ok {
        if &subResponse.Data != nil {
          o := subResponse.Data
            fmt.Printf("Order update, id: %d, state: %s, symbol: %s, price: %s, filled amount: %s", o.OrderId, o.OrderState, o.Symbol, o.Price, o.FilledAmount)
        }
      } else {
        fmt.Printf("Received unknown response: %v\n", resp)
      }
  })

// Connect to the server and wait for the handler to handle the response
err := client.Connect(true)
if err != nil {
  fmt.Printf("Client Connect error: %s\n", err)
  return
}
```

### Subscribe account change

*Authentication is required.*

```go
// Initialize a new instance
client := new(accountwebsocketclient.SubscribeAccountWebSocketV1Client).Init(config.AccessKey, config.SecretKey, config.Host)

// Set the callback handlers
client.SetHandler(
  // Connected handler
  func(resp *model.WebSocketV1AuthenticationResponse) {
    if resp.ErrorCode == 0 {
      err := client.Subscribe("1", "1250")
        if err != nil {
          fmt.Printf("Subscribe error: %s\n", err)
        } else {
          fmt.Println("Sent subscription")
        }
    } else {
      fmt.Printf("Authentication error: %d\n", resp.ErrorCode)
    }

  },
  // Response handler
  func(resp interface{}) {
    subResponse, ok := resp.(account.SubscribeAccountV1Response)
      if ok {
        if &subResponse.Data != nil {
          fmt.Printf("Account update event: %s\n", subResponse.Data.Event)
            if &subResponse.Data.List != nil {
              for _, b := range subResponse.Data.List {
                fmt.Printf("Account id: %d, currency: %s, type: %s, balance: %s", b.AccountId, b.Currency, b.Type, b.Balance)
              }
            }
        }
      } else {
        fmt.Printf("Received unknown response: %v\n", resp)
      }
  })

// Connect to the server and wait for the handler to handle the response
err := client.Connect(true)
if err != nil {
  fmt.Printf("Client Connect error: %s\n", err)
  return
}
```

## Unsubscribe

Since each websocket client manage the subscription separately, therefore you can cancel each individual subscription.