package helper

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/chinmayvivek/SwiftSchool/config"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// OpenDB initializes and returns a PostgreSQL database connection using the PostgresConfig.
func OpenPSQLDB(cfg *config.Config) (*sql.DB, error) {
	// Get a database connection using config's NewDBConnection method
	pSQLDB, err := cfg.NewPSQLDBConnection()
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return nil, err
	}
	SugarObj.Info("Successfully connected to the database.")
	return pSQLDB, nil
}

// QueryRow executes a query that is expected to return a single row.
func QueryRow(pSQLDB *sql.DB, query string, args ...interface{}) *sql.Row {
	return pSQLDB.QueryRow(query, args...)
}

// Query executes a query that is expected to return multiple rows.
func Query(pSQLDB *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	return pSQLDB.Query(query, args...)
}

// Exec executes an SQL statement that doesn't return rows (INSERT, UPDATE, DELETE, etc.)
func Exec(pSQLDB *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	return pSQLDB.Exec(query, args...)
}

// TransactionWrapper wraps a function with a database transaction.
func Transaction(pSQLDB *sql.DB, fn func(tx *sql.Tx) error) error {
	// Begin the transaction
	tx, err := pSQLDB.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}

	// Execute the function within the transaction
	if err := fn(tx); err != nil {
		// If error occurs, rollback the transaction and check for rollback error
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			SugarObj.Error("Error rolling back transaction: %v", rollbackErr)
		}
		return fmt.Errorf("transaction failed, rollback: %v", err)
	}

	// Commit the transaction if everything is successful
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}
