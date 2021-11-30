package usecase

import (
	"blockchain-trading/entity"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type ExchangeInteractor struct {
	ExchangeRepository ExchangeRepository
}

func (ei *ExchangeInteractor) ConfirmBalace() (balance []entity.Balance, err error) {
	balance, err = ei.ExchangeRepository.GetBalance()
	return
}

func (ei *ExchangeInteractor) ViewTicker(query map[string]string) (ticker *entity.Ticker, err error) {
	ticker, err = ei.ExchangeRepository.GetTicker(query)
	return
}

func (ei *ExchangeInteractor) ViewRealtimeTicker(symbol string, ch chan<- entity.Ticker) {
	ei.ExchangeRepository.GetRealTimeTicker(symbol, ch)
}

func (ei *ExchangeInteractor) OrderRegularly(orderData *entity.OrderData, interval, deadline time.Duration) {
	tk := time.NewTicker(interval)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-tk.C:
				respOrder, err := ei.ExchangeRepository.SendOrder(orderData)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(time.Now())
				spew.Dump(respOrder)
			}
		}
	}()
	time.Sleep(deadline)
	tk.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

func (di *ExchangeInteractor) ConfirmOHLC(query map[string]string) (ohlcs []entity.OHLC, err error) {
	ohlcs, err = di.ExchangeRepository.GetOHLC(query)
	return
}

func (di *ExchangeInteractor) ConfirmAllCurrencyInfomation() (currencies []entity.Currency, err error) {
	currencies, err = di.ExchangeRepository.GetAllCurrencyInfomation()
	return
}
