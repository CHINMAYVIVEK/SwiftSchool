package response

import (
	"encoding/json"
	"net/http"
	"time"
)

// StandardResponse is a standard structure for all API responses
type StandardResponse struct {
	Status     string      `json:"status"`         // "success" or "error"
	StatusCode int         `json:"statusCode"`     // HTTP status code (e.g., 200, 400, etc.)
	Message    string      `json:"message"`        // Human-readable message
	Data       interface{} `json:"data,omitempty"` // The actual data (could be any type)
	Timestamp  string      `json:"timestamp"`      // Timestamp of the response
}

// NewSuccessResponse creates a successful response with the provided data and message
func NewSuccessResponse(data interface{}, message string, statusCode int) *StandardResponse {
	return &StandardResponse{
		Status:     "success",
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Timestamp:  time.Now().Format(time.RFC3339), // Current timestamp
	}
}

// NewErrorResponse creates an error response with the provided message and error details
func NewErrorResponse(message string, statusCode int, err interface{}) *StandardResponse {
	return &StandardResponse{
		Status:     "error",
		StatusCode: statusCode,
		Message:    message,
		Data:       err,
		Timestamp:  time.Now().Format(time.RFC3339), // Current timestamp
	}
}

// RespondWithJSON writes the StandardResponse as a JSON response with a given HTTP status code
func RespondWithJSON(w http.ResponseWriter, statusCode int, response *StandardResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
