package repository

import (
	"blockchain-trading/api"
	"blockchain-trading/model/entity"
	"encoding/json"
)

type TickerService interface {
	GetTicker(productCode string) (*entity.Ticker, error)
}

type tickerService struct {
	ac api.APIClient
}

func NewTickerRepository(ac api.APIClient) TickerService {
	return &tickerService{ac}
}

func (ts *tickerService) GetTicker(productCode string) (*entity.Ticker, error) {
	url := "ticker"
	resp, err := ts.ac.DoRequest("GET", url, map[string]string{"product_code": productCode}, nil)
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
