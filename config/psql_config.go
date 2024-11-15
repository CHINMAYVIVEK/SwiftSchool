package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// PostgresConfig holds PostgreSQL connection configuration
type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port     int    `env:"POSTGRES_PORT" env-default:"5432"`
	Username string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Database string `env:"POSTGRES_DB" env-default:"postgres"`
	SSLMode  string `env:"POSTGRES_SSL_MODE" env-default:"disable"`
}

// NewPSQLDBConnection creates a PostgreSQL DB connection using the config
func (c *Config) NewPSQLDBConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.Postgres.Username,
		c.Postgres.Password,
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.Database,
		c.Postgres.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %v", err)
	}
	return db, nil
}
