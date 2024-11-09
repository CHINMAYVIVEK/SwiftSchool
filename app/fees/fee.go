package fees

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chinmayvivek/SwiftSchool/app/response"
	"github.com/chinmayvivek/SwiftSchool/config"
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

// Fee represents the fee data structure with enrollment and class information
type Fee struct {
	Head   string `json:"head"`
	Amount string `json:"amount"`
}

// FeeService manages the fee-related logic and database connection
type FeeService struct {
	DB *sql.DB // Database connection
}

// NewFeeService creates and initializes a FeeService
func NewFeeService(cfg *config.Config) (*FeeService, error) {
	db, err := helper.OpenDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize fee service: %v", err)
	}
	return &FeeService{DB: db}, nil

}

// GetAllFees returns all the fees in the system
func (fs *FeeService) GetFeeStructureByClass(w http.ResponseWriter, r *http.Request) {
	// Extract the "class_id" query parameter
	classID := r.URL.Query().Get("class_id")

	if classID == "" {
		// Missing class parameter, return 400 Bad Request
		response.RespondWithJSON(w, http.StatusBadRequest, response.NewErrorResponse("Class parameter is required", http.StatusBadRequest, nil))
		return
	}

	// Get the fee structure by class
	if res, err := getFeeStructures(classID, fs.DB); err != nil {
		// On error, return 500 Internal Server Error
		response.RespondWithJSON(w, http.StatusInternalServerError, response.NewErrorResponse(err.Error(), http.StatusInternalServerError, err))
	} else {
		// On success, return 200 OK with the result
		response.RespondWithJSON(w, http.StatusOK, response.NewSuccessResponse(res, "Fee structure retrieved successfully", http.StatusOK))
	}
}

func (fs *FeeService) UpdateFeeStructureByClass(w http.ResponseWriter, r *http.Request) {
	// Define the request structure with class_id and fee details
	var req []FeeByClass
	// Parse the request body and ensure it's valid
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || len(req) == 0 {
		response.RespondWithJSON(w, http.StatusBadRequest, response.NewErrorResponse("Invalid or missing payload", http.StatusBadRequest, err))
		return
	}

	// Pass the parsed data to updateFeeStructureByClass for processing
	if err := updateFeeStructureByClass(req, fs.DB); err != nil {
		response.RespondWithJSON(w, http.StatusInternalServerError, response.NewErrorResponse(err.Error(), http.StatusInternalServerError, err))
		return
	} else {
		// If all updates are successful, send a success response
		response.RespondWithJSON(w, http.StatusOK, response.NewSuccessResponse(nil, "Fee structure updated successfully", http.StatusOK))
	}

}

func (fs *FeeService) CreateNewFeeStructures(w http.ResponseWriter, r *http.Request) {

}
