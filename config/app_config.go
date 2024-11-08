package config

import "time"

type AppConfig struct {
	ReadHeaderTimeout time.Duration `env:"READ_HEADER_TIMEOUT" env-default:"30s"`
	LogLevel          string        `env:"LOG_LEVEL" env-default:"info"`
	ServerPort        int           `env:"SERVER_PORT" env-default:"8080"`
	ServerTimeout     time.Duration `env:"SERVER_TIMEOUT" env-default:"30s"`
}
