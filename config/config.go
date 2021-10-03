package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Key         string
	Secret      string
	ProductCode string
}

var Env EnvConfig

func init() {
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Env = EnvConfig{
		Key:         os.Getenv("API_KEY"),
		Secret:      os.Getenv("API_SECRET"),
		ProductCode: os.Getenv("PRODUCT_CODE"),
	}
}
