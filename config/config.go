package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	BfKey         string
	BfSecret      string
	CwKey         string
	CwSecret      string
	BinanceKey    string
	BinanceSecret string
	ProductCode   string
}

var Env EnvConfig

func init() {
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Env = EnvConfig{
		BfKey:         os.Getenv("BF_API_KEY"),
		BfSecret:      os.Getenv("BF_API_SECRET"),
		CwKey:         os.Getenv("CW_API_KEY"),
		CwSecret:      os.Getenv("CW_API_SECRET"),
		BinanceKey:    os.Getenv("BINANCE_API_KEY"),
		BinanceSecret: os.Getenv("BINANCE_API_SECRET"),
		ProductCode:   os.Getenv("PRODUCT_CODE"),
	}
}
