package presenter

import (
	"blockchain-trading/entity"
	"blockchain-trading/usecase"
	"context"
	"fmt"
)

type DatabasePresenter struct {
	Interactor *usecase.DatabaseInteractor
}

func (dp *DatabasePresenter) RegisterCurrencies(currencies []entity.Currency) error {
	for _, currency := range currencies {
		arg := usecase.ResisterCurrencyParams{
			Coin: currency.Coin,
			Name: currency.Name,
		}
		_, err := dp.Interactor.AddCurrency(context.Background(), arg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dp *DatabasePresenter) ShowCurrency(coin string) error {
	currency, err := dp.Interactor.FindCurrency(context.Background(), coin)
	if err != nil {
		return err
	}
	fmt.Println("id, coin, name")
	fmt.Printf("%d, %s, %s\n", currency.ID, currency.Coin, currency.Name)
	return nil
}

func (dp *DatabasePresenter) ShowCurrencies(arg usecase.ListCurrenciesParams) error {
	currencies, err := dp.Interactor.FindCurrencies(context.Background(), arg)
	if err != nil {
		return err
	}

	fmt.Println("id, coin, name")
	for _, currency := range currencies {
		fmt.Printf("%d, %s, %s\n", currency.ID, currency.Coin, currency.Name)
	}
	return nil
}
