package service

import (
	"blockchain-trading/model/entity"
	"blockchain-trading/model/repository"
	"encoding/json"
)

type TickerService interface {
	GetTicker(productCode string) (*entity.Ticker, error)
}

type tickerService struct {
	acr repository.APIClientRepository
}

func NewTickerService(acr repository.APIClientRepository) TickerService {
	return &tickerService{acr}
}

func (ts *tickerService) GetTicker(productCode string) (*entity.Ticker, error) {
	url := "ticker"
	resp, err := ts.acr.DoRequest("GET", url, map[string]string{"product_code": productCode}, nil)
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
