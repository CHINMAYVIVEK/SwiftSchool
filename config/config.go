package config

import (
	"context"
	"fmt"

	"github.com/caarlos0/env/v11"
)

// Config holds the entire application configuration
type Config struct {
	App      *AppConfig
	Postgres *PostgresDB
}

// New loads environment variables, initializes DB connections, and returns Config
func New(ctx context.Context) (*Config, error) {
	// Parse AppConfig
	appCfg := &AppConfig{}
	if err := env.Parse(appCfg); err != nil {
		return nil, fmt.Errorf("error parsing App config: %v", err)
	}

	// Parse Postgres config
	pgCfg := &PSQLConfig{}
	if err := env.Parse(pgCfg); err != nil {
		return nil, fmt.Errorf("error parsing PostgreSQL config: %v", err)
	}

	// Initialize PostgresDB
	pg := NewPostgresDB(pgCfg)
	if err := pg.Connect(ctx); err != nil {
		return nil, fmt.Errorf("error connecting to PostgreSQL: %v", err)
	}

	return &Config{
		App:      appCfg,
		Postgres: pg,
	}, nil
}

// Close safely closes all database connections
func (c *Config) Close(ctx context.Context) error {
	if err := c.Postgres.Close(); err != nil {
		return fmt.Errorf("error closing PostgreSQL connection: %v", err)
	}
	return nil
}
