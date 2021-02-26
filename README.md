# Huobi Golang SDK

This is Huobi Go SDK,  you can install to your Golang project and use this SDK to query all market data, trading and manage your account.

The SDK supports RESTful API invoking, and concurrently subscribing the market, account and order update from the Websocket connection.

## Table of Contents

- [Quick Start](#Quick-Start)
- [Usage](#Usage)
  - [Folder structure](#Folder-Structure)
  - [Run examples](#Run-examples)
  - [Client](#Client)
  - [Response](#Response)
  - [Init function](#Init-function)
  - [Logging](#logging)
- [Request examples](#Request-examples)
  - [Common data](#Common-data)
  - [Market data](#Market-data)
  - [Account](#account)
  - [Wallet](#wallet)
  - [Trading](#trading)
  - [Margin Loan](#margin-loan)
- [Subscription examples](#Subscription-examples)
  - [Subscribe candlestick update](#subscribe-candlestick-update)
  - [Subscribe account update](#subscribe-account-update)
  - [Subscribe order update](#subscribe-order-update)
  - [Subscribe trade update](#subscribe-trade-update)
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
  applogger.Error("Get timestamp error: %s", err)
} else {
  applogger.Info("Get timestamp: %d", resp)
}


// Get the list of accounts owned by this API user and print the detail on console
client := new(client.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host)
resp, err := client.GetAccountInfo()
if err != nil {
  applogger.Error("Get account error: %s", err)
} else {
  applogger.Info("Get account, count=%d", len(resp))
  for _, result := range resp {
    applogger.Info("account: %+v", result)
  }
}
```

## Usage

After above section, this SDK should be already download to your local machine, this section introduce this SDK and how to use it correctly.

### Folder Structure

This is the folder and package structure of SDK source code and the description

- **pkg**: The public package of the SDK
  - **client**: The client struct that are responsible to access data
  - **model**: The request and response data model
- **internal** The internal package that used internally
  - **gzip**: it provide the gzip decompress functionality that unzip the websocket binary data
  - **model**: The internal data model
  - **requestbuilder**: Responsible to build the request with the signature
- **logging**: It provides the logging function
- **config**: It stores the common configuration, such as host, access key.
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
|                 | CandlestickWebSocketClient        | Public  | WebSocket    |
|                 | DepthWebSocketClient              | Public  | WebSocket    |
|                 | MarketByPriceWebSocketClient      | Public  | WebSocket    |
|                 | BestBidOfferWebSocketClient       | Public  | WebSocket    |
|                 | TradeWebSocketClient              | Public  | WebSocket    |
|                 | Last24hCandlestickWebSocketClient | Public  | WebSocket    |
| Account         | AccountClient                     | Private | Rest         |
|                 | SubscribeAccountWebSocketV2Client | Private | WebSocket v2 |
| Wallet          | WalletClient                      | Private | Rest         |
| Order           | OrderClient                       | Private | Rest         |
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
  applogger.Error("Get timestamp error: %s", err)
} else {
  applogger.Info("Get timestamp: %d", resp)
}
```

For array struct, you can use for/range to iterate each element

```go
// Check the status of response and print some properties
for _, kline := range resp {
  applogger.Info("High=%v, Low=%v", kline.High, kline.Low)
}
```

### Init function

Golang is not a pure object oriented programming language, and there is no native constructor. In this SDK, every struct has an ***Init*** function for each struct, you must call Init function first, otherwise the member variables may not be initialized expected.

### Logging

This SDK uses the high performance logging library [zap](https://github.com/uber-go/zap), which provide different kind of loggers. To better support format message, this SDK uses the SugaredLogger, and wrapped a few interfaces in package *logging/applogger*. It has below features:

1. Logging target is console (In the future we will support output to file)
2. Support multiple levels (Fatal, Error, Panic, Warn, Info and Debug) and minimum log level
3. Support colorful text (by default)

You can customize your own logging by updating *applogger.go* file.

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
optionalRequest := market.GetCandlestickOptionalRequest{Period: market.MIN1, Size: 10}
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
optionalRequest := market.GetHistoricalTradeOptionalRequest{5}
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
queryDepositWithdrawOptionalRequest := wallet.QueryDepositWithdrawOptionalRequest{Currency: "usdt"}

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
  Price:     "1",
  Amount:    "10",
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
client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)  resp, err := client.GetOrderById("1")
```

#### Historical orders

```go
client := new(client.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host)
request := new(model.GetRequest).Init()
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
optionalRequest := margin.IsolatedMarginLoanOrdersOptionalRequest{
  StartDate: "2020-1-1",
}  

resp, err := client.MarginLoanOrders("btcusdt", optionalRequest)
```

## Subscription Examples

### Subscribe candlestick update

```go
// Initialize a new instance
client := new(marketwebsocketclient.Last24hCandlestickWebSocketClient).Init(config.Host)

// Set the callback handlers
client.SetHandler(
  // Connected handler
  func() {
    client.Request("btcusdt", "1608")

    client.Subscribe("btcusdt", "1608")
  },
  // Response handler
  func(resp interface{}) {
    candlestickResponse, ok := resp.(market.SubscribeLast24hCandlestickResponse)
      if ok {
        if &candlestickResponse != nil {
          if candlestickResponse.Tick != nil {
            t := candlestickResponse.Tick
              applogger.Info("Candlestick update, id: %d, count: %v, volume: %v [%v-%v-%v-%v]",
                         t.Id, t.Count, t.Vol, t.Open, t.Close, t.Low, t.High)
          }

          if candlestickResponse.Data != nil {
            t := candlestickResponse.Data
              applogger.Info("Candlestick data, id: %d, count: %v, volume: %v [%v-%v-%v-%v]",
                         t.Id, t.Count, t.Vol, t.Open, t.Close, t.Low, t.High)
          }
        }
      } else {
        applogger.Warn("Unknown response: %v", resp)
      }
  })

// Connect to the server and wait for the handler to handle the response
client.Connect(true)
```

### Subscribe account update

*Authentication is required.*

```go
// Initialize a new instance for account update websocket v2 client
  client := new(accountwebsocketclient.SubscribeAccountWebSocketV2Client).Init(config.AccessKey, config.SecretKey, config.Host)

  // Set the callback handlers
  client.SetHandler(
    // Authentication response handler
    func(resp *auth.WebSocketV2AuthenticationResponse) {
      if resp.IsSuccess() {
        client.Subscribe("1", "1149")        
      } else {
        applogger.Error("Authentication error, code: %d, message:%s", resp.Code, resp.Message)
      }
    },
    // Response handler
    func(resp interface{}) {
      subResponse, ok := resp.(account.SubscribeAccountV2Response)
      if ok {
        if subResponse.Action == "sub" {
          if subResponse.IsSuccess() {
            applogger.Info("Subscription topic %s successfully", subResponse.Ch)
          } else {
            applogger.Error("Subscription topic %s error, code: %d, message: %s", subResponse.Ch, subResponse.Code, subResponse.Message)
          }
        } else if subResponse.Action == "push" {
          if subResponse.Data != nil {
            b := subResponse.Data
            if b.ChangeTime == 0 {
              applogger.Info("Account overview, id: %d, currency: %s, balance: %s", b.AccountId, b.Currency, b.Balance)
            } else {
              applogger.Info("Account update, id: %d, currency: %s, balance: %s, time: %d", b.AccountId, b.Currency, b.Balance, b.ChangeTime)
            }
          }
        }
      } else {
        applogger.Warn("Received unknown response: %v", resp)
      }
    })

  // Connect to the server and wait for the handler to handle the response
  client.Connect(true)
```

### Subscribe order update

*Authentication is required.*

```go
// Initialize a new instance
  client := new(orderwebsocketclient.SubscribeOrderWebSocketV2Client).Init(config.AccessKey, config.SecretKey, config.Host)

  // Set the callback handlers
  client.SetHandler(
    // Connected handler
    func(resp *auth.WebSocketV2AuthenticationResponse) {
      if resp.IsSuccess() {
        // Subscribe if authentication passed
        client.Subscribe("btcusdt", "1149")
      } else {
        applogger.Info("Authentication error, code: %d, message:%s", resp.Code, resp.Message)
      }
    },
    // Response handler
    func(resp interface{}) {
      subResponse, ok := resp.(order.SubscribeOrderV2Response)
      if ok {
        if subResponse.Action == "sub" {
          if subResponse.IsSuccess() {
            applogger.Info("Subscription topic %s successfully", subResponse.Ch)
          } else {
            applogger.Error("Subscription topic %s error, code: %d, message: %s", subResponse.Ch, subResponse.Code, subResponse.Message)
          }
        } else if subResponse.Action == "push" {
          if subResponse.Data != nil {
            o := subResponse.Data
            applogger.Info("Order update, event: %s, symbol: %s, type: %s, status: %s",
              o.EventType, o.Symbol, o.Type, o.OrderStatus)
          }
        }
      } else {
        applogger.Warn("Received unknown response: %v", resp)
      }
    })

  // Connect to the server and wait for the handler to handle the response
  client.Connect(true)
```

### Subscribe trade update

*Authentication is required.*

```go
// Initialize a new instance
  client := new(orderwebsocketclient.SubscribeTradeClearWebSocketV2Client).Init(config.AccessKey, config.SecretKey, config.Host)

  // Set the callback handlers
  client.SetHandler(
    // Connected handler
    func(resp *auth.WebSocketV2AuthenticationResponse) {
      if resp.IsSuccess() {
        // Subscribe if authentication passed
        client.Subscribe("btcusdt", "1149")
      } else {
        applogger.Error("Authentication error, code: %d, message:%s", resp.Code, resp.Message)
      }
    },
    // Response handler
    func(resp interface{}) {
      subResponse, ok := resp.(order.SubscribeTradeClearResponse)
      if ok {
        if subResponse.Action == "sub" {
          if subResponse.IsSuccess() {
            applogger.Info("Subscription topic %s successfully", subResponse.Ch)
          } else {
            applogger.Error("Subscription topic %s error, code: %d, message: %s", subResponse.Ch, subResponse.Code, subResponse.Message)
          }
        } else if subResponse.Action == "push" {
          if subResponse.Data != nil {
            o := subResponse.Data
            applogger.Info("Order update, symbol: %s, order id: %d, price: %s, volume: %s",
              o.Symbol, o.OrderId, o.TradePrice, o.TradeVolume)
          }
        }
      } else {
        applogger.Warn("Received unknown response: %v", resp)
      }
    })

  // Connect to the server and wait for the handler to handle the response
  client.Connect(true)
```

## Unsubscribe

Since each websocket client manage the subscription separately, therefore you can cancel each individual subscription.