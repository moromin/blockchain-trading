package repository

import (
	"blockchain-trading/api"
	"blockchain-trading/model/entity"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type PastOHLCRepository interface {
	GetPastOHLC(params OHLCParams) (map[entity.Period][]entity.Interval, error)
}

type pastOHLCRepository struct {
	ac api.APIClient
}

func NewPastOHLCRepository(ac api.APIClient) PastOHLCRepository {
	return &pastOHLCRepository{ac}
}

type OHLCParams struct {
	ExchangeSymbol string
	PairSymbol     string
	Query          map[string]string
}

func (po *pastOHLCRepository) GetPastOHLC(params OHLCParams) (map[entity.Period][]entity.Interval, error) {
	urlPath := fmt.Sprintf("%s/%s/ohlc", params.ExchangeSymbol, params.PairSymbol)
	resp, err := po.ac.DoRequest("GET", urlPath, params.Query, nil, nil)
	if err != nil {
		return nil, err
	}

	var result entity.CryptoWatchResponse
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	} else if result.Error != "" {
		return nil, errors.New(result.Error)
	}

	srv := map[string][][]float64{}
	err = json.Unmarshal(result.Result, &srv)
	if err != nil {
		return nil, err
	}

	ret := make(map[entity.Period][]entity.Interval, len(srv))
	for srvPeriod, srvCandles := range srv {
		period := entity.Period(srvPeriod)
		candles := make([]entity.Interval, 0, len(srvCandles))
		for _, srvCandle := range srvCandles {
			if len(srvCandle) < 7 {
				return nil, fmt.Errorf("unexpected response from the server: wanted 7 elements, got %v", srvCandle)
			}
			ts := int64(srvCandle[0])
			candle := entity.Interval{
				Period:    period,
				CloseTime: time.Unix(ts, 0).UTC(),
				OHLC: entity.OHLC{
					Open:  srvCandle[1],
					High:  srvCandle[2],
					Low:   srvCandle[3],
					Close: srvCandle[4],
				},
				VolumeBase:  srvCandle[5],
				VolumeQuote: srvCandle[6],
			}
			candles = append(candles, candle)
		}
		ret[period] = candles
	}

	return ret, nil
}
