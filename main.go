package main

import (
	"blockchain-trading/config"
	"blockchain-trading/model/entity"
	"blockchain-trading/model/repository"
	"blockchain-trading/service"
	"fmt"
	"time"
)

func main() {
	fmt.Println(config.Env)

	apiClient := repository.NewAPIRepository(config.Env.Key, config.Env.Secret)

	balance := service.NewBalanceService(apiClient)
	fmt.Println(balance.GetBalance())

	// ticker := service.NewTickerService(apiClient)
	// fmt.Println(ticker.GetTicker(config.Env.ProductCode))

	tickerChannel := make(chan entity.Ticker)
	realTimeTicker := service.NewRealTimeTickerService(apiClient)
	go realTimeTicker.GetRealTimeTicker(config.Env.ProductCode, tickerChannel)
	for ticker := range tickerChannel {
		fmt.Println(ticker.GetMidPrice())
		fmt.Println(ticker.DateTime())
		fmt.Println(ticker.TruncateDateTime(time.Hour))
	}
}
