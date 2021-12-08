package usecase

import (
	"blockchain-trading/entity"
	"context"
)

type SqlcInteractor struct {
	Querier Querier
}

func (si *SqlcInteractor) AddCurrency(ctx context.Context, arg ResisterCurrencyParams) (entity.Currency, error) {
	currency, err := si.Querier.ResisterCurrency(ctx, arg)
	return currency, err
}

func (si *SqlcInteractor) FindCurrency(ctx context.Context, coin string) (entity.Currency, error) {
	currency, err := si.Querier.GetCurrency(context.Background(), coin)
	return currency, err
}

func (si *SqlcInteractor) FindCurrencies(ctx context.Context, arg ListCurrenciesParams) ([]entity.Currency, error) {
	currencies, err := si.Querier.ListCurrencies(context.Background(), arg)
	return currencies, err
}
