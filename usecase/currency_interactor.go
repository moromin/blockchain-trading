package usecase

import (
	"blockchain-trading/entity"
	"context"
)

type CurrencyInteractor struct {
	Repo CurrencyRepository
}

func (di *CurrencyInteractor) AddCurrency(ctx context.Context, arg ResisterCurrencyParams) (entity.Currency, error) {
	currency, err := di.Repo.RegisterCurrency(ctx, arg)
	return currency, err
}

func (di *CurrencyInteractor) FindCurrency(ctx context.Context, coin string) (entity.Currency, error) {
	currency, err := di.Repo.GetCurrency(context.Background(), coin)
	return currency, err
}

func (di *CurrencyInteractor) FindCurrencies(ctx context.Context, arg ListCurrenciesParams) ([]entity.Currency, error) {
	currencies, err := di.Repo.ListCurrencies(context.Background(), arg)
	return currencies, err
}
