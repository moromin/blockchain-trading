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
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func main() {
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

	startTime := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
	query := map[string]string{
		"symbol":    "BTCUSDT",
		"interval":  exchange.Interval1m,
		"startTime": strconv.Itoa(int(startTime.Unix() * 1000)),
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

	driverName, dataSourceName := config.SetDBConfig()
	conn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	container, err = di.NewOHLC(conn)
	if err != nil {
		panic(err)
	}

	if err := container.Invoke(func(op *presenter.OHLCPresenter) {
		err = op.RegisterOHLCs(ohlcs)
		if err != nil {
			panic(err)
		}
	}); err != nil {
		panic(err)
	}

}
