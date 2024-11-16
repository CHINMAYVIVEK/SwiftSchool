package config

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/chinmayvivek/SwiftSchool/helper"
)

// PostgresConfig holds the PostgreSQL database configuration
type PostgresConfig struct {
	Host            string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port            int    `env:"POSTGRES_PORT" env-default:"5432"`
	Username        string `env:"POSTGRES_USER"`
	Password        string `env:"POSTGRES_PASSWORD"`
	Database        string `env:"POSTGRES_DB" env-default:"postgres"`
	SSLMode         string `env:"POSTGRES_SSL_MODE" env-default:"disable"`
	MaxOpen         int    `env:"POSTGRES_MAX_OPEN" env-default:"10"`
	MaxIdle         int    `env:"POSTGRES_MAX_IDLE" env-default:"5"`
	ConnMaxLifetime int    `env:"POSTGRES_CONN_MAX_LIFETIME" env-default:"3600"`
}

// NewPSQLDBConnection creates a PostgreSQL DB connection with pooling based on the configuration
func (c *Config) NewPSQLDBConnection() (*sql.DB, error) {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.Postgres.Username, c.Postgres.Password, c.Postgres.Host, c.Postgres.Port, c.Postgres.Database, c.Postgres.SSLMode)

	// Open DB connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		helper.SugarObj.Error("Failed to open DB:", err)
		return nil, err
	}

	// Set connection pool options
	db.SetMaxOpenConns(c.Postgres.MaxOpen)
	db.SetMaxIdleConns(c.Postgres.MaxIdle)
	db.SetConnMaxLifetime(time.Duration(c.Postgres.ConnMaxLifetime) * time.Second)

	// Ping the DB to ensure connectivity
	if err := db.Ping(); err != nil {
		helper.SugarObj.Error("Failed to ping DB:", err)
		return nil, err
	}

	helper.SugarObj.Info("Successfully connected to the PostgreSQL database!")
	return db, nil
}
