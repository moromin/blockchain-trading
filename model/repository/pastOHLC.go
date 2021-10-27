package repository

import (
	"blockchain-trading/api"
	"blockchain-trading/model/entity"
	"encoding/json"
	"fmt"
	"time"
)

type PastOHLCRepository interface {
	GetPastOHLC(exchangeSymbol, pairSymbol string) (map[entity.Period][]entity.Interval, error)
}

type pastOHLCRepository struct {
	ac api.APIClient
}

func NewPastOHLCRepository(ac api.APIClient) PastOHLCRepository {
	return &pastOHLCRepository{ac}
}

func (po *pastOHLCRepository) GetPastOHLC(exchangeSymbol, pairSymbol string) (map[entity.Period][]entity.Interval, error) {
	urlPath := fmt.Sprintf("%s/%s/ohlc", exchangeSymbol, pairSymbol)
	resp, err := po.ac.DoRequest("GET", urlPath, map[string]string{}, nil, nil)
	if err != nil {
		return nil, err
	}

	var result entity.CryptoWatchResponse
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	} else if result.Error != "" {
		return nil, fmt.Errorf(result.Error)
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
