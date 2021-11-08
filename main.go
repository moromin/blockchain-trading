package main

import (
	"blockchain-trading/api"
	"blockchain-trading/config"
	"blockchain-trading/model/entity"
	"blockchain-trading/model/repository"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

const (
	dateFormat   = "2006-01-02 15:04:05"
	jsonFilePath = "json/"
)

func main() {
	// fmt.Println(config.Env)

	bfTarget := api.Target{
		BaseURL: api.BitFlyerURL,
		Header: map[string]string{
			"ACCESS-KEY":   config.Env.BfKey,
			"Content-Type": "application/json",
		},
	}
	bfClient := api.NewAPIClient(bfTarget)

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

	// cwParams := rest.RESTClientParams{
	// 	URL:    "", // If URL is empty, the default RESTURL will be specified.
	// 	APIKey: config.Env.CwKey,
	// }
	// cwClient := rest.NewRESTClient(&cwParams)
	// after, err := time.Parse(dateFormat, "2021-11-05 00:00:00")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// cwQuery := map[string]string{
	// 	"after":   strconv.Itoa(int(after.Unix())),
	// 	"periods": string(common.Period1H),
	// }
	// fmt.Println(cwQuery)
	// ohlc, err := cwClient.GetOHLC("bitflyer", "btcjpy", cwQuery)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%+v\n", ohlc)
	// order := repository.NewOrderRepository(bfClient)

	// Create order repository
	// Set order parameters (ex. "product_code": "BTC_JPY")
	order := repository.NewOrderRepository(bfClient)
	var orderParams entity.Order
	if err := orderParams.SetOrderParams(jsonFilePath + "order.json"); err != nil {
		fmt.Println(err)
		return
	}
	spew.Dump(orderParams)

	// Send order every 10 minutes for an hour.
	orderInterval := time.Minute * 10
	orderDeadline := time.Minute * 60
	tk := time.NewTicker(orderInterval)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-tk.C:
				respOrder, err := order.SendOrder(&orderParams)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(time.Now())
				spew.Dump(respOrder)
			}
		}
	}()
	time.Sleep(orderDeadline)
	tk.Stop()
	done <- true
	fmt.Println("Ticker stopped")

	// respOrder, err := order.SendOrder(&orderParams)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// spew.Dump(respOrder)
}
