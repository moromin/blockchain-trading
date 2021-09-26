package main

import (
	"blockchain-trading/bitflyer"
	"blockchain-trading/config"
	"fmt"
	"time"
)

func main() {
	apiClient := bitflyer.New(config.Env.Key, config.Env.Secret)
	ticker, _ := apiClient.GetTicker("BTC_JPY")
	fmt.Println(ticker)
	fmt.Println(ticker.GetMidPrice())
	fmt.Println(ticker.DateTime())
	fmt.Println(ticker.TruncateDateTime(time.Hour))
}
