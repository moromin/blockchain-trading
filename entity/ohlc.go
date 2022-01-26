package entity

import "time"

type OHLC struct {
	ID                       int64     `json:"-"`
	Symbol                   string    `json:"symbol"`
	Interval                 string    `json:"interval"`
	OpenTime                 time.Time `json:"time"`
	Open                     float64   `json:"open"`
	High                     float64   `json:"high"`
	Low                      float64   `json:"low"`
	Close                    float64   `json:"close"`
	Volume                   float64   `json:"volume"`
	CloseTime                time.Time `json:"-"`
	QuoteAssetVolume         float64   `json:"-"`
	NumberOfTrades           int64     `json:"-"`
	TakerBuyBaseAssetVolume  float64   `json:"-"`
	TakerBuyQuoteAssetVolume float64   `json:"-"`
}
