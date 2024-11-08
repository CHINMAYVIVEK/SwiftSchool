package fees

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"

	"github.com/chinmayvivek/SwiftSchool/helper"
)

type FeeStructureByClass struct {
	Data []FeeByClass `json:"data"`
}
type FeeByClass struct {
	ClassID   string `json:"class_id"`
	ClassName string `json:"class_name"`
	Fee       []Fee  `json:"fee"`
}

func updateFeeStructureByClass(req []FeeByClass, db *sql.DB) error {
	// Wrap the logic in a transaction
	return helper.Transaction(db, func(tx *sql.Tx) error {
		sqlQuery := `UPDATE fees SET amount = $1 WHERE class_id = $2 AND head = $3`
		var params []interface{}

		// Collect fee updates into the params slice
		for _, classData := range req {
			for _, fee := range classData.Fee {
				params = append(params, fee.Amount, classData.ClassID, fee.Head)
			}
		}

		// If there are updates, execute the query
		if len(params) > 0 {
			if _, err := tx.Exec(sqlQuery, params...); err != nil {
				helper.SugarObj.Error("Failed to update fee structure: %v", err)
				return fmt.Errorf("failed to update fee structure: %v", err)
			}
		} else {
			helper.SugarObj.Warn("No fee updates found")
		}
		return nil
	})
}

// Get Fee Structures  based on class IDs
func getFeeStructures(classID string, db *sql.DB) ([]FeeByClass, error) {
	// Split the classID string by `|~|` delimiter
	parts := strings.Split(classID, "|~|")
	classIDs := make([]string, 0, len(parts))     // Pre-allocate capacity
	placeholders := make([]string, 0, len(parts)) // Pre-allocate capacity

	// Iterate over the split parts
	for i, cid := range parts {
		// Trim spaces and validate in one step
		cleanedID := strings.TrimSpace(cid)
		if cleanedID == "" || !helper.IsValidID(cleanedID) {
			continue // Skip empty or invalid class IDs
		}

		// Append valid class ID and corresponding placeholder
		classIDs = append(classIDs, cleanedID)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}

	// If no valid class IDs are found, return nil
	if len(classIDs) == 0 {
		return nil, nil
	}

	// Prepare SQL queries
	sqlClassName := fmt.Sprintf("SELECT class_name FROM class WHERE class_id IN (%s)", strings.Join(placeholders, ","))
	sqlFees := fmt.Sprintf("SELECT head, amount FROM fees WHERE class_id IN (%s)", strings.Join(placeholders, ","))

	// Create channels to collect results
	classNamesChan := make(chan string, len(classIDs))
	feesChan := make(chan Fee, len(classIDs))
	errChan := make(chan error, 2)

	// Use WaitGroup to wait for both goroutines to complete
	var wg sync.WaitGroup

	// Fetch class names for the given class IDs
	wg.Add(1)
	go func() {
		defer wg.Done()
		rows, queryErr := helper.Query(db, sqlClassName, helper.ToInterfaceSlice(classIDs)...)
		if queryErr != nil {
			errChan <- fmt.Errorf("failed to fetch class names: %v", queryErr)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var className string
			if scanErr := rows.Scan(&className); scanErr != nil {
				errChan <- fmt.Errorf("failed to scan class name: %v", scanErr)
				return
			}
			classNamesChan <- className
		}
		if rows.Err() != nil {
			errChan <- fmt.Errorf("error during class name iteration: %v", rows.Err())
		}
		close(classNamesChan)
	}()

	// Fetch fee details for the given class IDs
	wg.Add(1)
	go func() {
		defer wg.Done()
		rows, queryErr := helper.Query(db, sqlFees, helper.ToInterfaceSlice(classIDs)...)
		if queryErr != nil {
			errChan <- fmt.Errorf("failed to query fee structure: %v", queryErr)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var fee Fee
			if scanErr := rows.Scan(&fee.Head, &fee.Amount); scanErr != nil {
				errChan <- fmt.Errorf("failed to scan fee row: %v", scanErr)
				return
			}
			feesChan <- fee
		}
		if rows.Err() != nil {
			errChan <- fmt.Errorf("error during fee iteration: %v", rows.Err())
		}
		close(feesChan)
	}()

	// Wait for both goroutines to complete
	wg.Wait()

	// Check for errors in the error channel
	select {
	case err := <-errChan:
		return nil, err
	default:
	}

	// Collect the results from the channels and combine them into FeeByClass
	var feeStructures []FeeByClass
	for className := range classNamesChan {
		fee := <-feesChan
		feeStructures = append(feeStructures, FeeByClass{
			// ClassID:   classID,
			ClassName: className,
			Fee:       []Fee{fee},
		})
	}

	// Return the combined result
	return feeStructures, nil
}
