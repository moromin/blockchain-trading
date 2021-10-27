package main

import (
	"blockchain-trading/api"
	"blockchain-trading/config"
	"blockchain-trading/model/repository"
	"fmt"
)

func main() {
	// fmt.Println(config.Env)

	// bfTarget := api.Target{
	// 	BaseURL: api.BitFlyerURL,
	// 	Header: map[string]string{
	// 		"ACCESS-KEY":   config.Env.BfKey,
	// 		"Content-Type": "application/json",
	// 	},
	// }
	// bfClient := api.NewAPIClient(bfTarget)

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

	cwTarget := api.Target{
		BaseURL: api.CryptoWatchURL,
		Header: map[string]string{
			"X-CW-API-Key": config.Env.CwKey,
			"User-Agent":   fmt.Sprintf("cw-sdk-go@%s", api.CwSdkVersion),
		},
	}
	cwAPIClient := api.NewAPIClient(cwTarget)
	res := repository.NewPastOHLCRepository(cwAPIClient)
	res.GetPastOHLC("bitflyer", "btcjpy")
}
