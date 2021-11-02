package main

import (
	"blockchain-trading/config"
	"fmt"

	"code.cryptowat.ch/cw-sdk-go/client/rest"
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

	cwParams := rest.RESTClientParams{
		URL:    "", // If URL is empty, the default RESTURL will be specified.
		APIKey: config.Env.CwKey,
	}
	cwClient := rest.NewRESTClient(&cwParams)
	ohlc, err := cwClient.GetOHLC("bitflyer", "btcjpy")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", ohlc)
}
