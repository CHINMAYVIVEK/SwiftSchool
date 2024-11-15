package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	App      *AppConfig
	Postgres *PostgresConfig
}

// NewConfig initializes and loads the config
func NewConfig() *Config {
	c := &Config{
		App:      &AppConfig{},
		Postgres: &PostgresConfig{},
	}
	c.Load()
	return c
}

// Load loads environment variables and configures the structs
func (c *Config) Load() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
	}

	// Parse environment variables into the Config struct
	if err := env.Parse(c); err != nil {
		panic(err)
	}
}
