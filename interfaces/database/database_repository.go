package database

import (
	"blockchain-trading/entity"
	"fmt"

	"github.com/pkg/errors"
)

type DatabaseRepository struct {
	SqlHandler SqlHandler
}

func (dr *DatabaseRepository) StoreCurrency(currencies []entity.Currency) error {
	for id, currency := range currencies {
		_, err := dr.SqlHandler.Execute("INSERT INTO currencies (coin, name) VALUES ($1, $2)", currency.Coin, currency.Name)
		if err != nil {
			fmt.Println(id, "failed ")
			return errors.Wrap(err, "Insert currency")
		}
	}
	return nil
}

func (dr *DatabaseRepository) FindAllCurrency() ([]entity.Currency, error) {
	rows, err := dr.SqlHandler.Query("SELECT * FROM currencies")
	defer func() {
		err = rows.Close()
		if err != nil {
			err = errors.Wrap(err, "Close row")
		}
	}()
	if err != nil {
		return nil, errors.Wrap(err, "Find all currency")
	}
	var currencies []entity.Currency
	for rows.Next() {
		var id int
		var coin string
		var name string
		if err := rows.Scan(&id, &coin, &name); err != nil {
			continue
		}
		currency := entity.Currency{
			ID:   id,
			Coin: coin,
			Name: name,
		}
		currencies = append(currencies, currency)
	}
	return currencies, err
}
