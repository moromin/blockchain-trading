package usecase

import (
	"blockchain-trading/entity"
	"context"
)

type DatabaseInteractor struct {
	DbRepo DatabaseRepository
}

func (di *DatabaseInteractor) AddCurrency(ctx context.Context, arg ResisterCurrencyParams) (entity.Currency, error) {
	currency, err := di.DbRepo.ResisterCurrency(ctx, arg)
	return currency, err
}

func (di *DatabaseInteractor) FindCurrency(ctx context.Context, coin string) (entity.Currency, error) {
	currency, err := di.DbRepo.GetCurrency(context.Background(), coin)
	return currency, err
}

func (di *DatabaseInteractor) FindCurrencies(ctx context.Context, arg ListCurrenciesParams) ([]entity.Currency, error) {
	currencies, err := di.DbRepo.ListCurrencies(context.Background(), arg)
	return currencies, err
}
