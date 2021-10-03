package repository

import (
	"blockchain-trading/api"
	"blockchain-trading/model/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

const (
	realtimeAPIScheme = "wss"
	realtimeAPIHost   = "ws.lightstream.bitflyer.com"
	realtimeAPIPath   = "/json-rpc"
	jsonRPCVersion    = "2.0"
)

type RealTimeTickerRepository interface {
	GetRealTimeTicker(symbol string, ch chan<- entity.Ticker)
}

type realTimeTickerRepository struct {
	ac api.APIClient
}

func NewRealTimeTickerRepository(ac api.APIClient) RealTimeTickerRepository {
	return &realTimeTickerRepository{ac}
}

func (rtts *realTimeTickerRepository) GetRealTimeTicker(symbol string, ch chan<- entity.Ticker) {
	u := url.URL{Scheme: realtimeAPIScheme, Host: realtimeAPIHost, Path: realtimeAPIPath}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	channel := fmt.Sprintf("lightning_ticker_%s", symbol)
	if err := c.WriteJSON(&entity.JsonRPC2{Version: jsonRPCVersion, Method: "subscribe", Params: &entity.SubscribeParams{channel}}); err != nil {
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
