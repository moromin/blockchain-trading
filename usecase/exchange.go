package usecase

import (
	"blockchain-trading/domain/model"
	"blockchain-trading/domain/repository"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type ExchangeUsecase interface {
	ConfirmBalace() ([]model.Balance, error)
	ViewTicker(productCode string) (*model.Ticker, error)
	ViewRealtimeTicker(symbol string, ch chan<- model.Ticker)
	OrderRegularly(orderData *model.OrderData, interval, deadline time.Duration)
}

type exchangeUsecase struct {
	exRepo repository.ExchangeRepository
}

func NewExchangeUsecase(exRepo repository.ExchangeRepository) ExchangeUsecase {
	return &exchangeUsecase{exRepo}
}

func (eu *exchangeUsecase) ConfirmBalace() (balance []model.Balance, err error) {
	balance, err = eu.exRepo.GetBalance()
	return
}

func (eu *exchangeUsecase) ViewTicker(productCode string) (ticker *model.Ticker, err error) {
	ticker, err = eu.exRepo.GetTicker(productCode)
	return
}

func (eu *exchangeUsecase) ViewRealtimeTicker(symbol string, ch chan<- model.Ticker) {
	eu.exRepo.GetRealTimeTicker(symbol, ch)
}

func (eu *exchangeUsecase) OrderRegularly(orderData *model.OrderData, interval, deadline time.Duration) {
	tk := time.NewTicker(interval)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-tk.C:
				respOrder, err := eu.exRepo.SendOrder(orderData)
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
