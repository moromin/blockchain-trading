package entity

type Currency struct {
	ID   int
	Coin string `json:"coin"`
	Name string `json:"name"`
}
