package fees

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/chinmayvivek/SwiftSchool/app/response"
)

type FeeStructureByClass struct {
	Data []FeeByClass `json:"data"`
}

type FeeByClass struct {
	ClassID   string `json:"class_id"`
	ClassName string `json:"class_name"`
	Fee       []Fee  `json:"fee"`
}

type Fee struct {
	Head   string `json:"head"`
	Amount string `json:"amount"`
}

type FeeService struct {
	DB *sql.DB // Database connection
}

// NewFeeService creates a new FeeService instance
func NewFeeService(db *sql.DB) *FeeService {
	return &FeeService{DB: db}
}

// GetFeeStructureByClass returns the fee structure for a given class
func (fs *FeeService) GetFeeStructureByClass(w http.ResponseWriter, r *http.Request) {
	classID := r.URL.Query().Get("class_id")
	if classID == "" {
		response.RespondWithJSON(w, http.StatusBadRequest, response.NewErrorResponse("Class parameter is required", http.StatusBadRequest, nil))
		return
	}

	if res, err := getFeeStructures(classID, fs.DB); err != nil {
		response.RespondWithJSON(w, http.StatusInternalServerError, response.NewErrorResponse(err.Error(), http.StatusInternalServerError, err))
	} else {
		response.RespondWithJSON(w, http.StatusOK, response.NewSuccessResponse(res, "Fee structure retrieved successfully", http.StatusOK))
	}
}

// UpdateFeeStructureByClass updates the fee structure for a given class
func (fs *FeeService) UpdateFeeStructureByClass(w http.ResponseWriter, r *http.Request) {
	var req []FeeByClass
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || len(req) == 0 {
		response.RespondWithJSON(w, http.StatusBadRequest, response.NewErrorResponse("Invalid or missing payload", http.StatusBadRequest, err))
		return
	}

	if err := updateFeeStructureByClass(req, fs.DB); err != nil {
		response.RespondWithJSON(w, http.StatusInternalServerError, response.NewErrorResponse(err.Error(), http.StatusInternalServerError, err))
	} else {
		response.RespondWithJSON(w, http.StatusOK, response.NewSuccessResponse(nil, "Fee structure updated successfully", http.StatusOK))
	}
}

// FeeCollection handles fee collection based on student enrollment and class
func (fs *FeeService) FeeCollection(w http.ResponseWriter, r *http.Request) {
	if res, err := feeCollection(); err != nil {
		response.RespondWithJSON(w, http.StatusInternalServerError, response.NewErrorResponse(err.Error(), http.StatusInternalServerError, err))
	} else {
		response.RespondWithJSON(w, http.StatusOK, response.NewSuccessResponse(res, "Fee collection retrieved successfully", http.StatusOK))
	}
}
