package database

import (
	"blockchain-trading/entity"
	"blockchain-trading/usecase"
	"context"
)

type CurrencyRepository struct {
	Db DBTX
}

const listCurrencies = `-- name: ListCurrencies :many
SELECT * FROM currencies
ORDER BY id
LIMIT $1
OFFSET $2
`

func (cr *CurrencyRepository) ListCurrencies(ctx context.Context, arg usecase.ListCurrenciesParams) ([]entity.Currency, error) {
	rows, err := cr.Db.QueryContext(ctx, listCurrencies, arg.Limit, arg.Offset)
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

const RegisterCurrency = `-- name: RegisterCurrency :one
INSERT INTO currencies (
  coin,
  name
) VALUES (
  $1, $2
) RETURNING *
`

func (cr *CurrencyRepository) RegisterCurrency(ctx context.Context, arg usecase.ResisterCurrencyParams) (entity.Currency, error) {
	row := cr.Db.QueryRowContext(ctx, RegisterCurrency, arg.Coin, arg.Name)
	var i entity.Currency
	err := row.Scan(&i.ID, &i.Coin, &i.Name)
	return i, err
}
