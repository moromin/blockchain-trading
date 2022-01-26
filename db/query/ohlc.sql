-- name: RegisterOHLC :one
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
) RETURNING *;

-- name: FindBySymbol :many
SELECT symbol, interval, opentime, open, high, low, close, volume
FROM ohlcs
WHERE symbol = $1
ORDER BY id;
