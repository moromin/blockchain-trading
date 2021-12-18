package currency

import (
	"blockchain-trading/entity"
	"blockchain-trading/interfaces/database"
	"blockchain-trading/usecase"
	"context"
)

type CurrencyRepository struct {
	Db database.DBTX
}

const getCurrency = `-- name: GetCurrency :one
SELECT * FROM currencies
WHERE coin = $1
`

func (dr *CurrencyRepository) GetCurrency(ctx context.Context, coin string) (entity.Currency, error) {
	row := dr.Db.QueryRowContext(ctx, getCurrency, coin)
	var i entity.Currency
	err := row.Scan(&i.ID, &i.Coin, &i.Name)
	return i, err
}

const listCurrencies = `-- name: ListCurrencies :many
SELECT * FROM currencies
ORDER BY id
LIMIT $1
OFFSET $2
`

func (dr *CurrencyRepository) ListCurrencies(ctx context.Context, arg usecase.ListCurrenciesParams) ([]entity.Currency, error) {
	rows, err := dr.Db.QueryContext(ctx, listCurrencies, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []entity.Currency{}
	for rows.Next() {
		var i entity.Currency
		if err := rows.Scan(&i.ID, &i.Coin, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const resisterCurrency = `-- name: ResisterCurrency :one
INSERT INTO currencies (
  coin,
  name
) VALUES (
  $1, $2
) RETURNING *
`

func (dr *CurrencyRepository) RegisterCurrency(ctx context.Context, arg usecase.ResisterCurrencyParams) (entity.Currency, error) {
	row := dr.Db.QueryRowContext(ctx, resisterCurrency, arg.Coin, arg.Name)
	var i entity.Currency
	err := row.Scan(&i.ID, &i.Coin, &i.Name)
	return i, err
}
