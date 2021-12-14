package presenter

import (
	"blockchain-trading/entity"
	"blockchain-trading/usecase"
	"context"
	"fmt"
)

type SqlcPresenter struct {
	Interactor *usecase.SqlcInteractor
}

func (sp *SqlcPresenter) RegisterCurrencies(currencies []entity.Currency) error {
	for _, currency := range currencies {
		arg := usecase.ResisterCurrencyParams{
			Coin: currency.Coin,
			Name: currency.Name,
		}
		_, err := sp.Interactor.AddCurrency(context.Background(), arg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sp *SqlcPresenter) ShowCurrency(coin string) error {
	currency, err := sp.Interactor.FindCurrency(context.Background(), coin)
	if err != nil {
		return err
	}
	fmt.Println("id, coin, name")
	fmt.Printf("%d, %s, %s\n", currency.ID, currency.Coin, currency.Name)
	return nil
}

func (sp *SqlcPresenter) ShowCurrencies(arg usecase.ListCurrenciesParams) error {
	currencies, err := sp.Interactor.FindCurrencies(context.Background(), arg)
	if err != nil {
		return err
	}

	fmt.Println("id, coin, name")
	for _, currency := range currencies {
		fmt.Printf("%d, %s, %s\n", currency.ID, currency.Coin, currency.Name)
	}
	return nil
}
