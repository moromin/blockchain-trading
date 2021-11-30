package entity

type Symbol struct {
	Symbols []struct {
		Symbol     string `json:"symbol"`
		BaseAsset  string `json:"baseAsset"`
		QuoteAsset string `json:"quoteAsset"`
	} `json:"symbols"`
}
