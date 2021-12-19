package usecase

import (
	"blockchain-trading/entity"
	"context"
)

type CurrencyInteractor struct {
	Repo CurrencyRepository
}

func (ci *CurrencyInteractor) AddCurrency(ctx context.Context, arg ResisterCurrencyParams) (entity.Currency, error) {
	currency, err := ci.Repo.RegisterCurrency(ctx, arg)
	return currency, err
}

func (ci *CurrencyInteractor) FindCurrencies(ctx context.Context, arg ListCurrenciesParams) ([]entity.Currency, error) {
	currencies, err := ci.Repo.ListCurrencies(context.Background(), arg)
	return currencies, err
}
