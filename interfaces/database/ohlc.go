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
