package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"swiftschool/internal/db"
)

// Database wraps the database connection and provides helper methods
type Database struct {
	conn          *sql.DB
	queryTimeout  time.Duration
	defaultUserID string // For created_by/updated_by fields
}

// New creates a new Database instance
func New(conn *sql.DB, queryTimeout time.Duration) *Database {
	return &Database{
		conn:         conn,
		queryTimeout: queryTimeout,
	}
}

// Queries returns a new SQLC Queries instance
func (d *Database) Queries() (*db.Queries, error) {
	if d.conn == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	return db.New(d.conn), nil
}

// WithTimeout creates a context with the configured query timeout
func (d *Database) WithTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, d.queryTimeout)
}

// BeginTx starts a new transaction
func (d *Database) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return d.conn.BeginTx(ctx, nil)
}

// QueriesWithTx returns a new SQLC Queries instance with a transaction
func (d *Database) QueriesWithTx(tx *sql.Tx) *db.Queries {
	return db.New(tx)
}

// GetConnection returns the underlying database connection
func (d *Database) GetConnection() *sql.DB {
	return d.conn
}

// SetDefaultUser sets the default user ID for audit fields
func (d *Database) SetDefaultUser(userID string) {
	d.defaultUserID = userID
}

// GetDefaultUser returns the default user ID
func (d *Database) GetDefaultUser() string {
	return d.defaultUserID
}
