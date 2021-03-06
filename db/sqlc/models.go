// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Currency struct {
	ID   int64  `json:"id"`
	Coin string `json:"coin"`
	Name string `json:"name"`
}

type Interval struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Ohlc struct {
	ID                       int64     `json:"id"`
	Symbol                   string    `json:"symbol"`
	Interval                 string    `json:"interval"`
	Opentime                 time.Time `json:"opentime"`
	Open                     string    `json:"open"`
	High                     string    `json:"high"`
	Low                      string    `json:"low"`
	Close                    string    `json:"close"`
	Volume                   string    `json:"volume"`
	Closetime                time.Time `json:"closetime"`
	QuoteAssetVolume         string    `json:"quote_asset_volume"`
	NumerOfTrades            int64     `json:"numer_of_trades"`
	TakerBuyBaseAssetVolume  string    `json:"taker_buy_base_asset_volume"`
	TakerBuyQuoteAssetVolume string    `json:"taker_buy_quote_asset_volume"`
}

type Symbol struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	BaseID  int64  `json:"base_id"`
	QuoteID int64  `json:"quote_id"`
}
