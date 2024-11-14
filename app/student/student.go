package student

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chinmayvivek/SwiftSchool/app/response"
	"github.com/chinmayvivek/SwiftSchool/config"
	"github.com/chinmayvivek/SwiftSchool/helper"
)

type Student struct {
	EnrollmentID string  `json:"enrollment_id"`
	Name         string  `json:"name"`
	ClassID      int     `json:"class_id"`
	Address      Address `json:"address"`
}

type Address struct {
	CurrentAddress   CurrentAddress   `json:"current_address"`
	PermanentAddress PermanentAddress `json:"permanent_address"`
	ContactDetails   ContactDetails   `json:"contact_details"`
}

type CurrentAddress struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

type PermanentAddress struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

type ContactDetails struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// FeeService manages the fee-related logic and database connection
type StudentService struct {
	DB *sql.DB // Database connection
}

// NewFeeService creates and initializes a FeeService
func NewStudentService(cfg *config.Config) (*StudentService, error) {
	db, err := helper.OpenPSQLDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize fee service: %v", err)
	}
	return &StudentService{DB: db}, nil

}

func (s *StudentService) RegisterStudent(w http.ResponseWriter, r *http.Request) {
	var req Student
	// Parse the request body and ensure it's valid
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.RespondWithJSON(w, http.StatusBadRequest, response.NewErrorResponse("Invalid or missing payload", http.StatusBadRequest, err))
		return
	}

	// register student
	if res, err := registerStudent(req, s.DB); err != nil {
		// On error, return 500 Internal Server Error
		response.RespondWithJSON(w, http.StatusInternalServerError, response.NewErrorResponse(err.Error(), http.StatusInternalServerError, err))
	} else {
		// On success, return 200 OK with the result
		response.RespondWithJSON(w, http.StatusOK, response.NewSuccessResponse(res, "Registration successful", http.StatusOK))
	}
}
