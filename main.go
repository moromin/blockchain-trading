package main

import (
	"blockchain-trading/config"
	"fmt"
)

func main() {
	fmt.Printf("%+v\n", config.Env)
	fmt.Println(config.Env.Key)
	fmt.Println(config.Env.Secret)
}
