package main

import (
	"blockchain-trading/config"
	"blockchain-trading/di"
	"blockchain-trading/entity"
	"blockchain-trading/infrastructure"
	"blockchain-trading/interfaces/exchange"
	"blockchain-trading/interfaces/presenter"
	"fmt"
	"net/url"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/lib/pq"
)

func main() {
	// Confirm apiclient
	// baseURL, err := url.Parse(exchange.BitFlyerURL)
	// if err != nil {
	// 	panic(err)
	// }

	// target := infrastructure.Target{
	// 	BaseURL: baseURL,
	// 	Header: map[string]string{
	// 		"ACCESS-KEY":   config.Env.BfKey,
	// 		"Content-Type": "application/json",
	// 	},
	// }

	// Setup binance API client
	baseURL, err := url.Parse(exchange.BinanceURL)
	if err != nil {
		panic(err)
	}

	target := infrastructure.Target{
		BaseURL: baseURL,
		Header: map[string]string{
			"X-MBX-APIKEY": config.Env.BinanceKey,
			"Content-Type": "application/json",
		},
	}

	container, err := di.NewAPIClient(target)
	if err != nil {
		fmt.Println(err)
		return
	}

	query := map[string]string{
		"symbol":   "BTCUSDT",
		"interval": exchange.Interval1m,
	}

	var ohlcs []entity.OHLC
	if err := container.Invoke(func(d *presenter.ExchangePresenter) {
		ohlcs, err = d.GetOHLC(query)
		if err != nil {
			panic(err)
		}
	}); err != nil {
		fmt.Println(err)
		return
	}

	spew.Dump(ohlcs)

	// driverName, dataSourceName := config.SetDBConfig()
	// conn, err := sql.Open(driverName, dataSourceName)
	// if err != nil {
	// 	panic(err)
	// }

	// container, err := di.NewCurrency(conn)
	// if err != nil {
	// 	panic(err)
	// }

	// if err := container.Invoke(func(dp *presenter.CurrencyPresenter) {

	// }); err != nil {
	// 	panic(err)
	// }

}
