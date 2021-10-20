package repository

import (
	"blockchain-trading/api"
	"blockchain-trading/model/entity"
	"encoding/json"
)

type TickerRepository interface {
	GetTicker(productCode string) (*entity.Ticker, error)
}

type tickerRepository struct {
	ac api.APIClient
}

func NewTickerRepository(ac api.APIClient) TickerRepository {
	return &tickerRepository{ac}
}

func (ts *tickerRepository) GetTicker(productCode string) (*entity.Ticker, error) {
	urlPath := "ticker"
	resp, err := ts.ac.DoRequest("GET", urlPath, map[string]string{"product_code": productCode}, nil, nil)
	if err != nil {
		return nil, err
	}

	var ticker entity.Ticker
	err = json.Unmarshal(resp, &ticker)
	if err != nil {
		return nil, err
	}
	return &ticker, nil
}
