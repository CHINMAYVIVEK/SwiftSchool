package helper

import (
	"context"
	"time"

	"swiftschool/config"
	"swiftschool/internal/db" // sqlc-generated package
)

// PostgresWrapper provides convenience methods for db access
type PostgresWrapper struct {
	postgres *config.PostgresDB
	timeout  time.Duration
}

// NewPostgresWrapper creates a wrapper
func NewPostgresWrapper(postgres *config.PostgresDB) *PostgresWrapper {
	return &PostgresWrapper{
		postgres: postgres,
		timeout:  postgres.QueryTimeout(),
	}
}

// WithTimeout returns a context with default timeout
func (w *PostgresWrapper) WithTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, w.timeout)
}

// Queries returns a sqlc Queries object
func (w *PostgresWrapper) Queries() (*db.Queries, error) {
	dbConn, err := w.postgres.GetDB()
	if err != nil {
		return nil, err
	}
	return db.New(dbConn), nil
}

// WithTx executes a function within a transaction
func (w *PostgresWrapper) WithTx(ctx context.Context, fn func(*db.Queries) error) error {
	dbConn, err := w.postgres.GetDB()
	if err != nil {
		return err
	}

	tx, err := dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queries := db.New(tx)
	if err := fn(queries); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
