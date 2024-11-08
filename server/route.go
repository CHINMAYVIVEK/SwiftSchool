package server

import (
	"net/http"

	"github.com/chinmayvivek/SwiftSchool/app/fees"
	"github.com/chinmayvivek/SwiftSchool/helper"
)

// LoadRoutes sets up the routes for the server
func (s *Server) LoadRoutes(mux *http.ServeMux) {

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
