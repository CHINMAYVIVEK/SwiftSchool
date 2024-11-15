package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// PostgresConfig holds PostgreSQL connection configuration
type PostgresConfig struct {
	Host            string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port            int    `env:"POSTGRES_PORT" env-default:"5432"`
	Username        string `env:"POSTGRES_USER"`
	Password        string `env:"POSTGRES_PASSWORD"`
	Database        string `env:"POSTGRES_DB" env-default:"postgres"`
	SSLMode         string `env:"POSTGRES_SSL_MODE" env-default:"disable"`
	MaxOpen         int    `env:"POSTGRES_MAX_OPEN" env-default:"10"`            // Max open connections
	MaxIdle         int    `env:"POSTGRES_MAX_IDLE" env-default:"5"`             // Max idle connections
	ConnMaxLifetime int    `env:"POSTGRES_CONN_MAX_LIFETIME" env-default:"3600"` // Max connection lifetime in seconds
}

// NewPSQLDBConnection creates a PostgreSQL DB connection using the config with pooling
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

	// Open DB connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %v", err)
	}

	// Configure connection pooling
	db.SetMaxOpenConns(c.Postgres.MaxOpen)                                         // Set max open connections
	db.SetMaxIdleConns(c.Postgres.MaxIdle)                                         // Set max idle connections
	db.SetConnMaxLifetime(time.Duration(c.Postgres.ConnMaxLifetime) * time.Second) // Set max connection lifetime

	// Check if the DB is reachable
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %v", err)
	}

	return db, nil
}
