package presenter

import (
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
