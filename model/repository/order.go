package repository

import (
	"blockchain-trading/api"
	"blockchain-trading/model/entity"
	"encoding/json"
)

type OrderRepository interface {
	SendOrder(order *entity.Order) (*entity.ResponseChildOrder, error)
}

type orderRepository struct {
	ac api.APIClient
}

func NewOrderRepository(ac api.APIClient) OrderRepository {
	return &orderRepository{ac}
}

func (or *orderRepository) SendOrder(order *entity.Order) (*entity.ResponseChildOrder, error) {
	method := "POST"
	urlPath := "me/sendchildorder"
	data, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	header := api.GetBitFlyerPrivateHeader(method, urlPath, data)
	resp, err := or.ac.DoRequest(method, urlPath, map[string]string{}, header, data)
	if err != nil {
		return nil, err
	}

	var respOrder entity.ResponseChildOrder
	err = json.Unmarshal(resp, &respOrder)
	if err != nil {
		return nil, err
	}
	return &respOrder, nil
}
