package main

import (
	"../logging"
	"./accountclientexample"
	"./accountwebsocketclientexample"
	"./commonclientexample"
	"./crossmarginclientexample"
	"./etfclientexample"
	"./isolatedmarginclientexample"
	"./marketclientexample"
	"./marketwebsocketclientexample"
	"./orderclientexample"
	"./orderwebsocketclientexample"
	"./walletclientexample"
)

func main() {
	runAll()
}

// Run all examples
func runAll() {
	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
	marketwebsocketclientexample.RunAllExamples()
	accountwebsocketclientexample.RunAllExamples()
	orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func perfTest() {
	logging.EnablePerformanceLog(true)

	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}
