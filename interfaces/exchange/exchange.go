package exchange

import (
	"blockchain-trading/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type ExchangeRepository struct {
	APIClient APIClient
}

func (er *ExchangeRepository) GetBalance() ([]entity.Balance, error) {
	endpoint := GetBalance
	method := endpoint.Method()
	urlPath := endpoint.String()
	header := endpoint.Header(nil)
	resp, err := er.APIClient.DoRequest(method, urlPath, nil, header, nil)
	if err != nil {
		return nil, err
	}

	var balance []entity.Balance
	err = json.Unmarshal(resp, &balance)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal balance")
	}
	return balance, nil
}

func (er *ExchangeRepository) GetTicker(query map[string]string) (*entity.Ticker, error) {
	endpoint := GetTicker
	method := endpoint.Method()
	urlPath := endpoint.String()
	header := endpoint.Header(nil)
	resp, err := er.APIClient.DoRequest(method, urlPath, query, header, nil)
	if err != nil {
		return nil, err
	}

	var ticker entity.Ticker
	err = json.Unmarshal(resp, &ticker)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal ticker")
	}
	return &ticker, nil
}

func (er *ExchangeRepository) GetRealTimeTicker(symbol string, ch chan<- entity.Ticker) {
	u := url.URL{Scheme: realtimeAPIScheme, Host: realtimeAPIHost, Path: realtimeAPIPath}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	channel := fmt.Sprintf("lightning_ticker_%s", symbol)
	if err := c.WriteJSON(&entity.JsonRPC2{Version: jsonRPCVersion, Method: "subscribe", Params: &entity.SubscribeParams{Channel: channel}}); err != nil {
		log.Fatal("subscribe:", err)
		return
	}

OUTER:
	for {
		message := new(entity.JsonRPC2)
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
				var ticker entity.Ticker
				if err := json.Unmarshal(marshaTic, &ticker); err != nil {
					continue OUTER
				}
				ch <- ticker
			}
		}
	}
}

func (er *ExchangeRepository) SendOrder(orderData *entity.OrderData) (*entity.Order, error) {
	data, err := json.Marshal(orderData)
	if err != nil {
		panic(err)
	}
	endpoint := SendOrder
	method := endpoint.Method()
	urlPath := endpoint.String()
	header := endpoint.Header(data)
	resp, err := er.APIClient.DoRequest(method, urlPath, nil, header, data)
	if err != nil {
		return nil, err
	}

	var order entity.Order
	err = json.Unmarshal(resp, &order)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal order")
	}
	return &order, nil
}
