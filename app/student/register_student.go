package student

import (
	"database/sql"
	"time"

	"github.com/chinmayvivek/SwiftSchool/helper"
	"github.com/google/uuid"
)

func registerStudent(req Student, db *sql.DB) (string, error) {
	// Generate UUID for student_id
	studentID := uuid.New().String()

	// Get the last two digits of the current year
	year := time.Now().Year() % 100

	// Generate the enrollment ID based on class and year
	classID := req.ClassID

	// SQL query to get the next serial number and insert the student
	sqlQuery := `
		WITH next_serial AS (
			SELECT COALESCE(MAX(CAST(SUBSTRING(enrollment_id FROM 5 FOR 4) AS INT)), 0) + 1 AS serial_number
			FROM student
			WHERE class_id = $1 AND EXTRACT(YEAR FROM created_at) = $2
		)
		INSERT INTO student (
			student_id, name, class_id,
			current_street, current_city, current_state, current_zip,
			permanent_street, permanent_city, permanent_state, permanent_zip,
			email, phone, enrollment_id
		)
		SELECT 
			$3, $4, $5, 
			$6, $7, $8, $9, 
			$10, $11, $12, $13, 
			$14, $15,
			-- Enrollment ID: class-year-serial_number
			CONCAT($5::TEXT, $2::TEXT, LPAD(next_serial.serial_number::TEXT, 4, '0'))
		FROM next_serial
		RETURNING enrollment_id
	`

	// Execute the query and retrieve the enrollment ID
	var enrollmentID string
	err := helper.QueryRow(db, sqlQuery,
		classID, year, studentID, req.Name, req.ClassID,
		req.Address.CurrentAddress.Street, req.Address.CurrentAddress.City, req.Address.CurrentAddress.State, req.Address.CurrentAddress.Zip,
		req.Address.PermanentAddress.Street, req.Address.PermanentAddress.City, req.Address.PermanentAddress.State, req.Address.PermanentAddress.Zip,
		req.Address.ContactDetails.Email, req.Address.ContactDetails.Phone,
	).Scan(&enrollmentID)

	if err != nil {
		helper.SugarObj.Error("Error registering student: %v", err)
		return "", err
	}

	// Return the generated enrollment ID
	return enrollmentID, nil
}
