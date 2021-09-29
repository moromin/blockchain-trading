package config

import (
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	Key         string
	Secret      string
	ProductCode string `envconfig:"PRODUCT_CODE" default:"BTC_JPY"`
}

var Env EnvConfig

func init() {
	envconfig.Process("api", &Env)
}
