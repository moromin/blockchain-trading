package main

import (
	"blockchain-trading/config"
	"blockchain-trading/model/repository"
	"blockchain-trading/service"
	"fmt"
)

func main() {
	fmt.Println(config.Env)

	apiClient := repository.NewAPIRepository(config.Env.Key, config.Env.Secret)

	balance := service.NewBalanceService(apiClient)
	fmt.Println(balance.GetBalance())

	ticker := service.NewTickerService(apiClient)
	fmt.Println(ticker.GetTicker(config.Env.ProductCode))
}
