package repository

import (
	"blockchain-trading/api"
	"blockchain-trading/model/entity"
	"encoding/json"
)

type CryptoWatchSampleRepository interface {
	GetMarketPrice() (*entity.CryptoWatchResponse, error)
}

type cryptoWatchSampleRepository struct {
	ac api.APIClient
}

func NewCryptoWatchSampleRepository(ac api.APIClient) CryptoWatchSampleRepository {
	return &cryptoWatchSampleRepository{ac}
}

func (cws *cryptoWatchSampleRepository) GetMarketPrice() (*entity.CryptoWatchResponse, error) {
	urlPath := "bitflyer/btcjpy/price"
	resp, err := cws.ac.DoRequest("GET", urlPath, map[string]string{}, nil, nil)
	if err != nil {
		return nil, err
	}

	var res entity.CryptoWatchResponse
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
