package presenter

import (
	"blockchain-trading/entity"
	"blockchain-trading/usecase"
	"context"
	"fmt"
)

type CurrencyPresenter struct {
	Interactor *usecase.CurrencyInteractor
}

func (cp *CurrencyPresenter) RegisterCurrencies(currencies []entity.Currency) error {
	for _, currency := range currencies {
		arg := usecase.ResisterCurrencyParams{
			Coin: currency.Coin,
			Name: currency.Name,
		}
		_, err := cp.Interactor.AddCurrency(context.Background(), arg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cp *CurrencyPresenter) ShowCurrency(coin string) error {
	currency, err := cp.Interactor.FindCurrency(context.Background(), coin)
	if err != nil {
		return err
	}
	fmt.Println("id, coin, name")
	fmt.Printf("%d, %s, %s\n", currency.ID, currency.Coin, currency.Name)
	return nil
}

func (cp *CurrencyPresenter) ShowCurrencies(arg usecase.ListCurrenciesParams) error {
	currencies, err := cp.Interactor.FindCurrencies(context.Background(), arg)
	if err != nil {
		return err
	}

	fmt.Println("id, coin, name")
	for _, currency := range currencies {
		fmt.Printf("%d, %s, %s\n", currency.ID, currency.Coin, currency.Name)
	}
	return nil
}
