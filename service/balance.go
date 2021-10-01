package service

import (
	"blockchain-trading/model/entity"
	"blockchain-trading/model/repository"
	"encoding/json"
	"log"
)

type BalanceService interface {
	GetBalance() ([]entity.Balance, error)
}

type balanceService struct {
	acr repository.APIClientRepository
}

func NewBalanceService(acr repository.APIClientRepository) BalanceService {
	return &balanceService{acr}
}

func (bs *balanceService) GetBalance() ([]entity.Balance, error) {
	url := "me/getbalance"
	resp, err := bs.acr.DoRequest("GET", url, map[string]string{}, nil)
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
