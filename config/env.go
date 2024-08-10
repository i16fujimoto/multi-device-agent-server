package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Env is 環境変数
type Env struct {
	AppEnv            string `envconfig:"APP_ENV" default:"local"`
	BasicAuthPassword string `envconfig:"BASIC_AUTH_PASSWORD" default:"secret"`
}

var env Env

func init() {
	if err := envconfig.Process("", &env); err != nil {
		panic(fmt.Errorf("failed to get environment variables: %w", err))
	}
}

func GetEnv() *Env {
	return &env
}

// IsLocal : Check if the current environment is local
func IsLocal() bool {
	return env.AppEnv == "local"
}

// IsTest : Check if the current environment is test
func IsTest() bool {
	return env.AppEnv == "test"
}

// IsDev : Check if the current environment is development
func IsDev() bool {
	return env.AppEnv == "dev"
}

// IsStg : Check if the current environment is staging
func IsStg() bool {
	return env.AppEnv == "stg"
}

// IsPrd : Check if the current environment is production
func IsPrd() bool {
	return env.AppEnv == "prd"
}
