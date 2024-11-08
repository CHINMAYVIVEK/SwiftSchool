package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// PostgresConfig contains the configuration for PostgreSQL database connection
type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port     int    `env:"POSTGRES_PORT" env-default:"5432"`
	Username string `env:"POSTGRES_USER" env-default:""`
	Password string `env:"POSTGRES_PASSWORD" env-default:""`
	Database string `env:"POSTGRES_DB" env-default:"postgres"`
	SSLMode  string `env:"POSTGRES_SSL_MODE" env-default:"disable"`
}

// NewDBConnection returns a PostgreSQL database connection using the PostgresConfig.
func (c *Config) NewDBConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.Postgres.Username,
		c.Postgres.Password,
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.Database,
		c.Postgres.SSLMode,
	)
	// Open the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgreSQL database: %v", err)
	}

	// Check if the database is reachable
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	return db, nil
}
