package entity

import "encoding/json"

type CryptoWatchResponse struct {
	Result json.RawMessage `json:"result"`
	Error  string          `json:"error"`
}
