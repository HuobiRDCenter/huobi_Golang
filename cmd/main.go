package main

import (
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

func perfTest() {
	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}
