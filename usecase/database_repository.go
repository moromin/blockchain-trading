package usecase

import "blockchain-trading/entity"

type DatabaseRepository interface {
	StoreCurrency(currencies []entity.Currency) error
	FindAllCurrency() ([]entity.Currency, error)
}
