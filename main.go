package main

import (
	"blockchain-trading/config"
	"fmt"
	"strconv"
	"time"

	"code.cryptowat.ch/cw-sdk-go/client/rest"
	"code.cryptowat.ch/cw-sdk-go/common"
)

const dateFormat = "2006-01-02 15:04:05"

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
	after, err := time.Parse(dateFormat, "2021-11-05 00:00:00")
	if err != nil {
		fmt.Println(err)
		return
	}
	cwQuery := map[string]string{
		"after":   strconv.Itoa(int(after.Unix())),
		"periods": string(common.Period1H),
	}
	// fmt.Println(cwQuery)
	ohlc, err := cwClient.GetOHLC("bitflyer", "btcjpy", cwQuery)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", ohlc)
}
