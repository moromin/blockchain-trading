package database

import (
	"blockchain-trading/entity"
	"blockchain-trading/usecase"
	"context"
)

type OHLCRepository struct {
	Db DBTX
}

const RegisterOHLC = `-- name: RegisterOHLC :one
INSERT INTO ohlcs (
	symbol,
	interval,
	opentime,
	open,
	high,
	low,
	close,
	volume,
	closetime,
	quote_asset_volume,
	number_of_trades,
	taker_buy_base_asset_volume,
	taker_buy_quote_asset_volume
) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
) RETURNING *
`

func (or *OHLCRepository) RegisterOHLC(ctx context.Context, arg usecase.RegisterOHLCParams) (entity.OHLC, error) {
	row := or.Db.QueryRowContext(ctx, RegisterOHLC,
		arg.Symbol,
		arg.Interval,
		arg.Opentime,
		arg.Open,
		arg.High,
		arg.Low,
		arg.Close,
		arg.Volume,
		arg.Closetime,
		arg.QuoteAssetVolume,
		arg.NumberOfTrades,
		arg.TakerBuyBaseAssetVolume,
		arg.TakerBuyQuoteAssetVolume,
	)
	var i entity.OHLC
	err := row.Scan(
		&i.ID,
		&i.Symbol,
		&i.Interval,
		&i.OpenTime,
		&i.Open,
		&i.High,
		&i.Low,
		&i.Close,
		&i.Volume,
		&i.CloseTime,
		&i.QuoteAssetVolume,
		&i.NumberOfTrades,
		&i.TakerBuyBaseAssetVolume,
		&i.TakerBuyQuoteAssetVolume,
	)
	return i, err
}

const FindBySymbol = `-- name: FindBySymbol :many
SELECT symbol, interval, opentime, open, high, low, close, volume
FROM ohlcs
WHERE symbol = $1
ORDER BY id
`

func (or *OHLCRepository) FindBySymbol(ctx context.Context, symbol string) ([]entity.OHLC, error) {
	rows, err := or.Db.QueryContext(ctx, FindBySymbol, symbol)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []entity.OHLC{}
	for rows.Next() {
		var i entity.OHLC
		if err := rows.Scan(
			&i.Symbol,
			&i.Interval,
			&i.OpenTime,
			&i.Open,
			&i.High,
			&i.Low,
			&i.Close,
			&i.Volume,
		); err != nil {
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
