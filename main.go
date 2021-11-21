package main

import (
	"blockchain-trading/config"
	"blockchain-trading/di"
	"blockchain-trading/infrastructure"
	"blockchain-trading/interfaces/exchange"
	"blockchain-trading/interfaces/presenter"
	"fmt"
	"net/url"
)

// Use this when call Cryptowatch SDK methods to get past OHLC data.
const dateFormat = "2006-01-02 15:04:05"

func main() {
	baseURL, err := url.Parse(exchange.BitFlyerURL)
	if err != nil {
		panic(err)
	}

	target := infrastructure.Target{
		BaseURL: baseURL,
		Header: map[string]string{
			"ACCESS-KEY":   config.Env.BfKey,
			"Content-Type": "application/json",
		},
	}

	container, err := di.New(target)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := container.Invoke(func(e *presenter.ExchangePresenter) {
		e.ShowBalance()
	}); err != nil {
		fmt.Println(err)
		return
	}
}
