package main

import (
	"blockchain-trading/config"
	"blockchain-trading/infra"
	"blockchain-trading/usecase"

	"github.com/davecgh/go-spew/spew"
)

// Use this when call Cryptowatch SDK methods to get past OHLC data.
const dateFormat = "2006-01-02 15:04:05"

func main() {
	target := infra.Target{
		BaseURL: infra.BitFlyerURL,
		Header: map[string]string{
			"ACCESS-KEY":   config.Env.BfKey,
			"Content-Type": "application/json",
		},
	}

	ac := infra.NewAPIClient(target)
	er := infra.NewExchangeRepository(ac)
	es := usecase.NewExchangeUsecase(er)

	balance, _ := es.ConfirmBalace()
	spew.Dump(balance)
}
