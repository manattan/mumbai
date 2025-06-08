package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppName      string `envconfig:"APP_NAME" default:"mumbai-service"`
	HTTPPort     int    `envconfig:"HTTP_PORT" default:"8080"`
	GRPCPort     int    `envconfig:"GRPC_PORT" default:"50051"`
	MySQLDSN     string `envconfig:"MYSQL_DSN" required:"true"`
	LogLevel     string `envconfig:"LOG_LEVEL" default:"info"`
	Env          string `envconfig:"ENV" default:"development"`
}

func Load() *Config {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("failed to load env: %v", err)
	}
	return &cfg
}