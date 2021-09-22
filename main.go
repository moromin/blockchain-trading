package main

import (
	"blockchain-trading/config"
	"fmt"
)

func main() {
	fmt.Println("ApiKey:", config.Config.ApiKey)
	fmt.Println("ApiSecret:", config.Config.ApiSecret)
}
