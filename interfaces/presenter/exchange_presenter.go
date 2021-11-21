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
