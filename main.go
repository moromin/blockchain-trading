package main

import (
	"blockchain-trading/bitflyer"
	"blockchain-trading/config"
	"fmt"
)

func main() {
	apiClient := bitflyer.New(config.Env.Key, config.Env.Secret)
	fmt.Println(apiClient.GetBalance())
}
