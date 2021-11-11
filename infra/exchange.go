package infra

import (
	"blockchain-trading/domain/model"
	"blockchain-trading/domain/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type exchangeRepository struct {
	ac APIClient
}

func NewExchangeRepository(ac APIClient) repository.ExchangeRepository {
	return &exchangeRepository{ac}
}

func (er *exchangeRepository) GetBalance() ([]model.Balance, error) {
	method := "GET"
	urlPath := "me/getbalance"
	header := getBitFlyerPrivateHeader(method, urlPath, nil)
	resp, err := er.ac.DoRequest(method, urlPath, map[string]string{}, header, nil)
	if err != nil {
		return nil, err
	}

	var balance []model.Balance
	err = json.Unmarshal(resp, &balance)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal balance")
	}
	return balance, nil
}

func (er *exchangeRepository) GetTicker(productCode string) (*model.Ticker, error) {
	urlPath := "ticker"
	resp, err := er.ac.DoRequest("GET", urlPath, map[string]string{"product_code": productCode}, nil, nil)
	if err != nil {
		return nil, err
	}

	var ticker model.Ticker
	err = json.Unmarshal(resp, &ticker)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal ticker")
	}
	return &ticker, nil
}

func (er *exchangeRepository) GetRealTimeTicker(symbol string, ch chan<- model.Ticker) {
	u := url.URL{Scheme: realtimeAPIScheme, Host: realtimeAPIHost, Path: realtimeAPIPath}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	channel := fmt.Sprintf("lightning_ticker_%s", symbol)
	if err := c.WriteJSON(&model.JsonRPC2{Version: jsonRPCVersion, Method: "subscribe", Params: &model.SubscribeParams{Channel: channel}}); err != nil {
		log.Fatal("subscribe:", err)
		return
	}

OUTER:
	for {
		message := new(model.JsonRPC2)
		if err := c.ReadJSON(message); err != nil {
			log.Println("read:", err)
			return
		}

		if message.Method == "channelMessage" {
			switch v := message.Params.(type) {
			case map[string]interface{}:
				marshaTic, err := json.Marshal(v["message"])
				if err != nil {
					continue OUTER
				}
				var ticker model.Ticker
				if err := json.Unmarshal(marshaTic, &ticker); err != nil {
					continue OUTER
				}
				ch <- ticker
			}
		}
	}
}

func (er *exchangeRepository) SendOrder(orderData *model.OrderData) (*model.Order, error) {
	method := "POST"
	urlPath := "me/sendchildorder"
	data, err := json.Marshal(orderData)
	if err != nil {
		return nil, errors.Wrap(err, "Marshal order data")
	}
	header := getBitFlyerPrivateHeader(method, urlPath, data)
	resp, err := er.ac.DoRequest(method, urlPath, map[string]string{}, header, data)
	if err != nil {
		return nil, err
	}

	var order model.Order
	err = json.Unmarshal(resp, &order)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal order")
	}
	return &order, nil
}
