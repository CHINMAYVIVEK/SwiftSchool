package helper

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// QueryRow executes a query that returns a single row.
func QueryRow(pSQLDB *sql.DB, query string, args ...interface{}) *sql.Row {
	return pSQLDB.QueryRow(query, args...)
}

// Query executes a query that returns multiple rows.
func Query(pSQLDB *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	return pSQLDB.Query(query, args...)
}

// Exec executes an SQL statement without returning rows (e.g., INSERT, UPDATE, DELETE).
func Exec(pSQLDB *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	return pSQLDB.Exec(query, args...)
}

// Transaction wraps a function with a database transaction.
func Transaction(pSQLDB *sql.DB, fn func(tx *sql.Tx) error) error {
	tx, err := pSQLDB.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}

	if err := fn(tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			SugarObj.Error("Error rolling back transaction: %v", rollbackErr)
		}
		return fmt.Errorf("transaction failed, rollback: %v", err)
	}

	return tx.Commit()
}
