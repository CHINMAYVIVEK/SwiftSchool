package helper

import (
	"context"
	"time"

	"swiftschool/config"
	"swiftschool/internal/db" // sqlc-generated package
)

// PostgresWrapper wraps DB operations with timeout + SQLC access
type PostgresWrapper struct {
	dbConn  *config.PostgresDB
	timeout time.Duration
}

// NewPostgresWrapper initializes the wrapper
func NewPostgresWrapper(dbConn *config.PostgresDB) *PostgresWrapper {
	return &PostgresWrapper{
		dbConn:  dbConn,
		timeout: dbConn.QueryTimeout(),
	}
}

// WithTimeout provides a context with default query timeout
func (w *PostgresWrapper) WithTimeout(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, w.timeout)
}

// Queries returns a SQLC Queries instance for normal DB operations
func (w *PostgresWrapper) Queries() (*db.Queries, error) {
	conn, err := w.dbConn.GetDB()
	if err != nil {
		return nil, err
	}
	return db.New(conn), nil
}

// WithTx runs a SQLC transaction wrapper
func (w *PostgresWrapper) WithTx(ctx context.Context, fn func(*db.Queries) error) error {
	conn, err := w.dbConn.GetDB()
	if err != nil {
		return err
	}

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)

	if err := fn(q); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
