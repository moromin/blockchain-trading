package main

import (
	"blockchain-trading/injector"

	"github.com/davecgh/go-spew/spew"
)

// Use this when call Cryptowatch SDK methods to get past OHLC data.
const dateFormat = "2006-01-02 15:04:05"

func main() {
	exchangeUsecase := injector.InjectExchangeUsecase()

	balance, _ := exchangeUsecase.ConfirmBalace()
	spew.Dump(balance)

	ticker, _ := exchangeUsecase.ViewTicker("BTC_JPY")
	spew.Dump(ticker)
}
