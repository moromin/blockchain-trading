package repository

import (
	"blockchain-trading/api"
	"blockchain-trading/model/entity"
	"encoding/json"
	"log"
)

type BalanceRepository interface {
	GetBalance() ([]entity.Balance, error)
}

type balanceRepository struct {
	ac api.APIClient
}

func NewBalanceRepository(ac api.APIClient) BalanceRepository {
	return &balanceRepository{ac}
}

func (bs *balanceRepository) GetBalance() ([]entity.Balance, error) {
	url := "me/getbalance"
	resp, err := bs.ac.DoRequest("GET", url, map[string]string{}, nil)
	log.Printf("url=%s resp=%s", url, string(resp))
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}

	var balance []entity.Balance
	err = json.Unmarshal(resp, &balance)
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	return balance, nil
}
