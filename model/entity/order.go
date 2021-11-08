package entity

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

type Order struct {
	ProductCode    string  `json:"product_code"`
	ChildOrderType string  `json:"child_order_type"`
	Side           string  `json:"side"`
	Price          int     `json:"price"`
	Size           float64 `json:"size"`
	MinuteToExpire int     `json:"minute_to_expire"`
	TimeInForce    string  `json:"time_in_force"`
}

type ResponseChildOrder struct {
	ChildOrderAcceptanceID string `json:"child_order_acceptance_id"`
}

func (o *Order) SetOrderParams(filepath string) error {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		return errors.Wrap(err, "read error")
	}
	if err = json.Unmarshal(raw, o); err != nil {
		return errors.Wrap(err, "json unmarshal error")
	}
	return nil
}
