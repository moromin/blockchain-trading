package presenter

import (
	"blockchain-trading/entity"
	"blockchain-trading/usecase"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type ExchangePresenter struct {
	Interactor *usecase.ExchangeInteractor
}

func (ep *ExchangePresenter) ShowBalance() {
	balance, err := ep.Interactor.ConfirmBalace()
	if err != nil {
		fmt.Println(err)
		return
	}
	spew.Dump(balance)
}

func (dp *ExchangePresenter) ShowOHLC(query map[string]string) {
	ohlcs, err := dp.Interactor.ConfirmOHLC(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	spew.Dump(ohlcs)
}

func (dp *ExchangePresenter) GetOHLC(query map[string]string) ([]entity.OHLC, error) {
	ohlcs, err := dp.Interactor.ConfirmOHLC(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return ohlcs, nil
}

func (dp *ExchangePresenter) GetAllCurrency() ([]entity.Currency, error) {
	currencies, err := dp.Interactor.ConfirmAllCurrencyInfomation()
	if err != nil {
		return nil, err
	}
	return currencies, nil
}
