package main

import (
	"blockchain-trading/api"
	"blockchain-trading/config"
	"blockchain-trading/model/entity"
	"blockchain-trading/model/repository"
	"fmt"
	"time"
)

func main() {
	fmt.Println(config.Env)

	apiClient := api.NewAPIClient(config.Env.Key, config.Env.Secret)

	balance := repository.NewBalanceRepository(apiClient)
	fmt.Println(balance.GetBalance())

	ticker := repository.NewTickerRepository(apiClient)
	fmt.Println(ticker.GetTicker(config.Env.ProductCode))

	tickerChannel := make(chan entity.Ticker)
	realTimeTicker := repository.NewRealTimeTickerRepository(apiClient)
	go realTimeTicker.GetRealTimeTicker(config.Env.ProductCode, tickerChannel)
	for ticker := range tickerChannel {
		fmt.Println(ticker.GetMidPrice())
		fmt.Println(ticker.DateTime())
		fmt.Println(ticker.TruncateDateTime(time.Hour))
	}
}
