package usecase

import (
	"blockchain-trading/entity"
	"context"
	"time"
)

type RegisterOHLCParams struct {
	Symbol                   string    `json:"symbol"`
	Interval                 string    `json:"interval"`
	Opentime                 time.Time `json:"opentime"`
	Open                     float64   `json:"open"`
	High                     float64   `json:"high"`
	Low                      float64   `json:"low"`
	Close                    float64   `json:"close"`
	Volume                   float64   `json:"volume"`
	Closetime                time.Time `json:"closetime"`
	QuoteAssetVolume         float64   `json:"quote_asset_volume"`
	NumberOfTrades           int64     `json:"number_of_trades"`
	TakerBuyBaseAssetVolume  float64   `json:"taker_buy_base_asset_volume"`
	TakerBuyQuoteAssetVolume float64   `json:"taker_buy_quote_asset_volume"`
}

type OHLCRepository interface {
	RegisterOHLC(ctx context.Context, arg RegisterOHLCParams) (entity.OHLC, error)
}
