package server

import (
	"encoding/json"
	"net/http"

	"github.com/chinmayvivek/SwiftSchool/app/fees"
	"github.com/chinmayvivek/SwiftSchool/helper"
)

// LoadRoutes sets up the routes for the server
func (s *Server) LoadRoutes(mux *http.ServeMux) {

	// Health check route - returns JSON response without using a struct
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// Set the response content type to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write status code and response body directly using a map
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]string{"status": "OK"}); err != nil {
			// Log the error if JSON encoding fails
			helper.SugarObj.Error("Failed to encode response: %v", err)
		}
	})

	// Initialize FeeService with the config
	feeService, err := fees.NewFeeService(s.Config)
	if err != nil {
		helper.SugarObj.Error("Failed to initialize fee service: %v", err)
		return
	}
	// Handle /api/fees/fee-structure-by-class for both GET and POST
	mux.HandleFunc("/api/fees/fee-structure-by-class", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Call the handler for GET method
			feeService.GetFeeStructureByClass(w, r)
		case http.MethodPost:
			// Call the handler for POST method
			feeService.UpdateFeeStructureByClass(w, r)
		default:
			// Handle unsupported HTTP methods
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
