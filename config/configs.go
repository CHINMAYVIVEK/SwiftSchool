package config

import (
	"github.com/caarlos0/env"
	"github.com/chinmayvivek/SwiftSchool/helper"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type Config struct {
	App      *AppConfig
	Postgres *PostgresConfig
}

// NewConfig initializes the configuration by loading environment variables
func NewConfig() *Config {
	c := &Config{
		App:      &AppConfig{},
		Postgres: &PostgresConfig{},
	}
	c.Load()
	return c
}

// Load loads environment variables from the .env file and parses them into the config struct
func (c *Config) Load() {
	if err := godotenv.Load(".env"); err != nil {
		helper.SugarObj.Error("Error loading .env file: ", err)
	}

	if err := env.Parse(c); err != nil {
		helper.SugarObj.Error("Error parsing environment variables: ", err)
	}
}
