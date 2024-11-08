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

// NewConfig initializes the config and loads environment variables.
func NewConfig() *Config {
	c := &Config{}
	c.Load()
	return c
}

// Load loads environment variables and configures the structs
func (c *Config) Load() {
	// Load environment variables from the .env file (if it exists)
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	c.App = &AppConfig{}
	c.Postgres = &PostgresConfig{}

	// Parse the environment variables into the Config struct
	if err := env.Parse(c); err != nil {
		panic(err)
	}
}
