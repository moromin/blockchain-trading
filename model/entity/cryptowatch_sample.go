package entity

type CryptoWatchSample struct {
	Result struct {
		Price float64 `json:"price"`
	} `json:"result"`
}
