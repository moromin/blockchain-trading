package presenter

import (
	"blockchain-trading/entity"
	"blockchain-trading/usecase"

	"github.com/davecgh/go-spew/spew"
)

type DatabasePresenter struct {
	Interactor *usecase.DatabaseInteractor
}

func (dp *DatabasePresenter) RegisterCurrencies(currencies []entity.Currency) error {
	err := dp.Interactor.Add(currencies)
	if err != nil {
		return err
	}
	return nil
}

func (dp *DatabasePresenter) ShowCurrencies() error {
	currencies, err := dp.Interactor.Currencies()
	if err != nil {
		return err
	}
	spew.Dump(currencies)
	return nil
}
