package main

import (
	"blockchain-trading/config"
	"blockchain-trading/di"
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

	// container, err := di.New(target)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if err := container.Invoke(func(e *presenter.ExchangePresenter) {
	// 	e.ShowBalance()
	// }); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// Confirm ohlc data
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

	container, err := di.New(target)
	if err != nil {
		fmt.Println(err)
		return
	}

	// query := map[string]string{
	// 	"symbol":   "BTCUSDT",
	// 	"interval": exchange.Interval1m,
	// }

	// var coins []entity.Coin
	if err := container.Invoke(func(d *presenter.ExchangePresenter) {
		coins, err := d.GetAllCoin()
		if err != nil {
			fmt.Println(err)
			return
		}
		spew.Dump(coins)
	}); err != nil {
		fmt.Println(err)
		return
	}

	// Confirm DB
	// conn, err := sql.Open("postgres", "user=root password=secret host=localhost dbname=ohlc sslmode=disable")
	// if err != nil {
	// 	panic(err)
	// }
	// handler := infrastructure.SqlHandler{Conn: conn}
	// container, err = di.NewDB(handler)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// symbol := "BTCUSDT"
	// if err := container.Invoke(func(dbp *presenter.DatabasePresenter) {
	// 	dbp.RegisterSymbol(symbol)
	// }); err != nil {
	// 	panic(err)
	// }
}
