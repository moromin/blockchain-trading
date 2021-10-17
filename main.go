package main

import (
	"blockchain-trading/api"
	"blockchain-trading/config"
	"blockchain-trading/model/repository"
	"fmt"
)

func main() {
	// fmt.Println(config.Env)

	// bfClient := api.NewAPIClient(config.Env.BfKey, config.Env.BfSecret, "bitflyer")

	// balance := repository.NewBalanceRepository(bfClient)
	// fmt.Println(balance.GetBalance())

	// ticker := repository.NewTickerRepository(bfClient)
	// fmt.Println(ticker.GetTicker(config.Env.ProductCode))

	// tickerChannel := make(chan entity.Ticker)
	// realTimeTicker := repository.NewRealTimeTickerRepository(bfClient)
	// go realTimeTicker.GetRealTimeTicker(config.Env.ProductCode, tickerChannel)
	// for ticker := range tickerChannel {
	// 	fmt.Println(ticker.GetMidPrice())
	// 	fmt.Println(ticker.DateTime())
	// 	fmt.Println(ticker.TruncateDateTime(time.Hour))
	// }

	cwAPIClient := api.NewAPIClient(config.Env.CwKey, config.Env.CwSecret, "cryptowatch")
	res := repository.NewCryptoWatchSampleRepository(cwAPIClient)
	fmt.Println(res.GetMarketPrice())
}
