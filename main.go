package main

import (
	"blockchain-trading/config"
	"fmt"
)

func main() {
	fmt.Println(config.Env)
	// apiClient := repository.NewAPIRepository(config.Env.Key, config.Env.Secret)

	// balance := service.NewBalanceService(apiClient)
	// fmt.Println(balance.GetBalance())
}
