package repository

import "blockchain-trading/domain/model"

type ExchangeRepository interface {
	GetBalance() ([]model.Balance, error)
	GetTicker(productCode string) (*model.Ticker, error)
	GetRealTimeTicker(symbol string, ch chan<- model.Ticker)
	SendOrder(orderData *model.OrderData) (*model.Order, error)
}
