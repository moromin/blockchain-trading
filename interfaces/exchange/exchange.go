package exchange

import (
	"blockchain-trading/entity"
	"blockchain-trading/interfaces/api"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type ExchangeRepository struct {
	APIClient api.APIClient
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

func (dr *ExchangeRepository) GetOHLC(query map[string]string) ([]entity.OHLC, error) {
	endpoint := GetOHLC
	method := endpoint.Method()
	urlPath := endpoint.String()
	header := endpoint.Header(nil)
	resp, err := dr.APIClient.DoRequest(method, urlPath, query, header, nil)
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		panic(err)
	}
	ohlcArray, ok := data.([]interface{})
	if !ok {
		panic(ok)
	}

	ohlcs := make([]entity.OHLC, len(ohlcArray))
	for i, ohlc := range ohlcArray {
		data, ok := ohlc.([]interface{})
		if !ok {
			panic(ok)
		}

		ohlcs[i].OpenTime = int64(data[0].(float64))
		ohlcs[i].Open, _ = strconv.ParseFloat(data[1].(string), 64)
		ohlcs[i].High, _ = strconv.ParseFloat(data[2].(string), 64)
		ohlcs[i].Low, _ = strconv.ParseFloat(data[3].(string), 64)
		ohlcs[i].Close, _ = strconv.ParseFloat(data[4].(string), 64)
		ohlcs[i].Volume, _ = strconv.ParseFloat(data[5].(string), 64)
		ohlcs[i].CloseTime = int64(data[6].(float64))
		ohlcs[i].QuoteAssetVolume, _ = strconv.ParseFloat(data[7].(string), 64)
		ohlcs[i].NumberOfTrades = int64(data[8].(float64))
		ohlcs[i].TakerBuyBaseAssetVolume, _ = strconv.ParseFloat(data[9].(string), 64)
		ohlcs[i].TakerBuyQuoteAssetVolume, _ = strconv.ParseFloat(data[10].(string), 64)
	}

	return ohlcs, nil
}
