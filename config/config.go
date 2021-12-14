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
	DBDriver      string
	DBUser        string
	DBPassword    string
	DBHost        string
	DBName        string
	DBSSLMode     string
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
		DBDriver:      os.Getenv("DBDriver"),
		DBUser:        os.Getenv("DBUser"),
		DBPassword:    os.Getenv("DBPassword"),
		DBHost:        os.Getenv("DBHost"),
		DBName:        os.Getenv("DBName"),
		DBSSLMode:     os.Getenv("DBSSLMode"),
	}
}

func SetDBConfig() (driverName, dataSourceName string) {
	driverName = Env.DBDriver
	dataSourceName = fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s", Env.DBUser, Env.DBPassword, Env.DBHost, Env.DBName, Env.DBSSLMode)
	return
}
