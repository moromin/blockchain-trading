package main

import (
	"blockchain-trading/bitflyer"
	"blockchain-trading/config"
	"fmt"
	"time"
)

func main() {
	apiClient := bitflyer.New(config.Env.Key, config.Env.Secret)

	fmt.Println(apiClient.GetBalance())
	tickerChannel := make(chan bitflyer.Ticker)
	go apiClient.GetRealTimeTicker(config.Env.ProductCode, tickerChannel)
	for ticker := range tickerChannel {
		fmt.Println(ticker.GetMidPrice())
		fmt.Println(ticker.DateTime())
		fmt.Println(ticker.TruncateDateTime(time.Hour))
	}
}
