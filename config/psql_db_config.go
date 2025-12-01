package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// PostgresDB manages the PostgreSQL connection
type PostgresDB struct {
	config *PSQLConfig
	db     *sql.DB
}

// PSQLConfig holds PostgreSQL database configuration
type PSQLConfig struct {
	// Either provide URL or individual fields
	URL      string `env:"POSTGRES_URL" default:""`
	Host     string `env:"POSTGRES_HOST" default:"localhost"`
	Port     string `env:"POSTGRES_PORT" default:"5432"`
	User     string `env:"POSTGRES_USER" default:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" default:"postgres"`
	DBName   string `env:"POSTGRES_DB_NAME" default:"swift_school"`

	SetMaxOpenConns    int           `env:"POSTGRES_MAX_OPEN_CONNS" default:"25"`
	SetMaxIdleConns    int           `env:"POSTGRES_MAX_IDLE_CONNS" default:"5"`
	SetConnMaxLifetime time.Duration `env:"POSTGRES_CONN_MAX_LIFETIME" default:"5m"`
	QueryTimeout       time.Duration `env:"POSTGRES_QUERY_TIMEOUT" default:"10s"`
}

// NewPostgresDB creates a new PostgresDB instance
func NewPostgresDB(config *PSQLConfig) *PostgresDB {
	return &PostgresDB{config: config}
}

// Connect establishes the connection to PostgreSQL
func (p *PostgresDB) Connect(ctx context.Context) error {
	dsn := p.buildDSN()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(p.config.SetMaxOpenConns)
	db.SetMaxIdleConns(p.config.SetMaxIdleConns)
	db.SetConnMaxLifetime(p.config.SetConnMaxLifetime)

	// Test connection using configured QueryTimeout
	timeout := p.QueryTimeout()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	p.db = db
	return nil
}

// buildDSN constructs the DSN from URL or individual fields
func (p *PostgresDB) buildDSN() string {
	if p.config.URL != "" {
		return p.config.URL
	}
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.config.Host, p.config.Port, p.config.User, p.config.Password, p.config.DBName,
	)
}

// GetDB returns the underlying *sql.DB
func (p *PostgresDB) GetDB() (*sql.DB, error) {
	if p.db == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	return p.db, nil
}

// Close safely closes the database connection
func (p *PostgresDB) Close() error {
	if p.db != nil {
		err := p.db.Close()
		p.db = nil
		return err
	}
	return nil
}

// QueryTimeout returns the configured query timeout or default
func (p *PostgresDB) QueryTimeout() time.Duration {
	if p.config.QueryTimeout == 0 {
		return 10 * time.Second
	}
	return p.config.QueryTimeout
}
