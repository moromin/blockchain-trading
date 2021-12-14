package usecase

import "blockchain-trading/entity"

type DatabaseInteractor struct {
	DatabaseRepository DatabaseRepository
}

func (di *DatabaseInteractor) Add(currencies []entity.Currency) error {
	err := di.DatabaseRepository.StoreCurrency(currencies)
	if err != nil {
		return err
	}
	return nil
}

func (di *DatabaseInteractor) Currencies() ([]entity.Currency, error) {
	currencies, err := di.DatabaseRepository.FindAllCurrency()
	if err != nil {
		return nil, err
	}
	return currencies, nil
}
