package fees

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"

	"github.com/chinmayvivek/SwiftSchool/helper"
)

// Update fee structure based on class and fee details
func updateFeeStructureByClass(req []FeeByClass, db *sql.DB) error {
	if len(req) == 0 {
		helper.SugarObj.Warn("No fee structure data provided.")
		return nil
	}

	// Estimate total parameters required for SQL query
	params := []interface{}{}
	for _, classData := range req {
		for _, fee := range classData.Fee {
			params = append(params, fee.Amount, classData.ClassID, fee.Head)
		}
	}

	if len(params) == 0 {
		helper.SugarObj.Warn("No fee updates found.")
		return nil
	}

	return helper.Transaction(db, func(tx *sql.Tx) error {
		sqlQuery := `UPDATE fees SET amount = $1 WHERE class_id = $2 AND head = $3`
		if _, err := tx.Exec(sqlQuery, params...); err != nil {
			helper.SugarObj.Error("Failed to update fee structure: %v", err)
			return fmt.Errorf("failed to update fee structure: %v", err)
		}
		helper.SugarObj.Info("Fee structure updated successfully.")
		return nil
	})
}

// Get fee structures based on class IDs
func getFeeStructures(classID string, db *sql.DB) ([]FeeByClass, error) {
	classIDs := []string{}
	placeholders := []string{}
	parts := strings.Split(classID, "|~|")

	for i, cid := range parts {
		cleanedID := strings.TrimSpace(cid)
		if cleanedID == "" || !helper.IsValidID(cleanedID) {
			continue
		}
		classIDs = append(classIDs, cleanedID)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}

	if len(classIDs) == 0 {
		helper.SugarObj.Error("No valid class IDs found")
		return nil, nil
	}

	// SQL queries for fetching class names and fees
	sqlClassName := fmt.Sprintf("SELECT class_name FROM class WHERE class_id IN (%s)", strings.Join(placeholders, ","))
	sqlFees := fmt.Sprintf("SELECT head, amount FROM fees WHERE class_id IN (%s)", strings.Join(placeholders, ","))

	helper.SugarObj.Info("Fetching class names and fee structures")

	// Channels and WaitGroup for concurrent execution
	classNamesChan := make(chan string, len(classIDs))
	feesChan := make(chan Fee, len(classIDs))
	errChan := make(chan error, 2)
	var wg sync.WaitGroup

	// Fetch class names concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		rows, err := helper.Query(db, sqlClassName, helper.ToInterfaceSlice(classIDs)...)
		if err != nil {
			errChan <- fmt.Errorf("failed to fetch class names: %v", err)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var className string
			if err := rows.Scan(&className); err != nil {
				errChan <- fmt.Errorf("failed to scan class name: %v", err)
				return
			}
			classNamesChan <- className
		}
		close(classNamesChan)
	}()

	// Fetch fee details concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		rows, err := helper.Query(db, sqlFees, helper.ToInterfaceSlice(classIDs)...)
		if err != nil {
			errChan <- fmt.Errorf("failed to query fee structure: %v", err)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var fee Fee
			if err := rows.Scan(&fee.Head, &fee.Amount); err != nil {
				errChan <- fmt.Errorf("failed to scan fee row: %v", err)
				return
			}
			feesChan <- fee
		}
		close(feesChan)
	}()

	// Wait for goroutines to complete
	wg.Wait()

	// Check for any errors
	select {
	case err := <-errChan:
		helper.SugarObj.Error("Error: %v", err)
		return nil, err
	default:
	}

	// Combine results from both channels
	var feeStructures []FeeByClass
	for className := range classNamesChan {
		fee := <-feesChan
		feeStructures = append(feeStructures, FeeByClass{
			ClassName: className,
			Fee:       []Fee{fee},
		})
	}

	helper.SugarObj.Info("Fee structures retrieved successfully.")
	return feeStructures, nil
}
