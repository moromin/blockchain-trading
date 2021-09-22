package config

import (
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	Key    string
	Secret string
}

var Env EnvConfig

func init() {
	envconfig.Process("api", &Env)
}
