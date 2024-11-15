package server

import (
	"encoding/json"
	"net/http"

	"github.com/chinmayvivek/SwiftSchool/app/fees"
	"github.com/chinmayvivek/SwiftSchool/app/student"
	"github.com/chinmayvivek/SwiftSchool/helper"
)

// LoadRoutes sets up the server's routes
func (s *Server) LoadRoutes(mux *http.ServeMux) {
	feeService := fees.NewFeeService(s.DB)
	studentService := student.NewStudentService(s.DB)

	// Health check route
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]string{"status": "OK"}); err != nil {
			helper.SugarObj.Error("Failed to encode response: %v", err)
		}
	})

	// Fee structure route (GET and POST)
	mux.HandleFunc("/api/fees/fee-structure-by-class", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			feeService.GetFeeStructureByClass(w, r)
		case http.MethodPost:
			feeService.UpdateFeeStructureByClass(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Student registration route (POST)
	mux.HandleFunc("/api/student/student-registration", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			studentService.RegisterStudent(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
