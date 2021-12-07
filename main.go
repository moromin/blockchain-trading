package main

import (
	"blockchain-trading/config"
	"blockchain-trading/di"
	"blockchain-trading/entity"
	"blockchain-trading/infrastructure"
	"blockchain-trading/interfaces/exchange"
	"blockchain-trading/interfaces/presenter"
	"database/sql"
	"fmt"
	"net/url"

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

	var currencies []entity.Currency
	if err := container.Invoke(func(d *presenter.ExchangePresenter) {
		currencies, err = d.GetAllCurrency()
		if err != nil {
			fmt.Println(err)
			return
		}
	}); err != nil {
		fmt.Println(err)
		return
	}

	// Show all currencies from binance API
	fmt.Println("id, coin, name")
	for i, data := range currencies {
		fmt.Println(i, data.Coin, data.Name)
	}

	// Confirm DB
	driverName, dataSourceName := config.SetDBConfig()
	conn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	handler := infrastructure.SqlHandler{Conn: conn}
	container, err = di.NewDB(handler)
	if err != nil {
		fmt.Println(err)
	}

	if err := container.Invoke(func(dbp *presenter.DatabasePresenter) {
		err = dbp.RegisterCurrencies(currencies)
		if err != nil {
			fmt.Println(err)
			return
		}
	}); err != nil {
		panic(err)
	}

}
