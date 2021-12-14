package usecase

import (
	"blockchain-trading/entity"
	"context"
)

type ListCurrenciesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ResisterCurrencyParams struct {
	Coin string `json:"coin"`
	Name string `json:"name"`
}

type DatabaseRepository interface {
	GetCurrency(ctx context.Context, coin string) (entity.Currency, error)
	ListCurrencies(ctx context.Context, arg ListCurrenciesParams) ([]entity.Currency, error)
	ResisterCurrency(ctx context.Context, arg ResisterCurrencyParams) (entity.Currency, error)
}
