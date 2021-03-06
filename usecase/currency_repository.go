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

type CurrencyRepository interface {
	ListCurrencies(ctx context.Context, arg ListCurrenciesParams) ([]entity.Currency, error)
	RegisterCurrency(ctx context.Context, arg ResisterCurrencyParams) (entity.Currency, error)
}
