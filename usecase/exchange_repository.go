package usecase

import "blockchain-trading/entity"

type ExchangeRepository interface {
	GetBalance() ([]entity.Balance, error)
	GetTicker(query map[string]string) (*entity.Ticker, error)
	GetRealTimeTicker(symbol string, ch chan<- entity.Ticker)
	SendOrder(orderData *entity.OrderData) (*entity.Order, error)
	GetOHLC(query map[string]string) ([]entity.OHLC, error)
	GetAllCurrencyInfomation() ([]entity.Currency, error)
}
