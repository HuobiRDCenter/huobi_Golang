package main

import (
	"github.com/huobirdcenter/huobi_golang/cmd/accountclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/accountwebsocketclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/algoorderclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/commonclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/crossmarginclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/etfclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/isolatedmarginclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/marketclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/marketwebsocketclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/orderclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/orderwebsocketclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/stablecoinclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/subuserclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/walletclientexample"
	"github.com/huobirdcenter/huobi_golang/logging/perflogger"
)

func main() {
	//subUserClient := new(client.SubUserClient).Init("eba1b128-67692151-edc99346-mjlpdje3ld", "2a54d214-5addbd2d-d715ef58-a349f", "api.huobi.pro")
	//id, err := subUserClient.GetUid()
	//if err != nil {
	//	fmt.Printf("get uid error:%s", err)
	//}
	//balance, err := subUserClient.GetSubUserAggregateBalance()
	//if err != nil {
	//	fmt.Printf("get balance error:%s", err)
	//}
	//fmt.Printf("balence:%s", balance)
	//getResp := "{\n  \"code\": 200,\n  \"data\": 63628520\n}"
	//result := account.GetUidResponse{}
	//jsonErr := json.Unmarshal([]byte(getResp), &result)
	//runAll()
}

// Run all examples
func runAll() {
	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	algoorderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	subuserclientexample.RunAllExamples()
	stablecoinclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
	marketwebsocketclientexample.RunAllExamples()
	accountwebsocketclientexample.RunAllExamples()
	orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)

	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}
