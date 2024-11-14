package server

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/chinmayvivek/SwiftSchool/app/fees"
	"github.com/chinmayvivek/SwiftSchool/app/student"
	"github.com/chinmayvivek/SwiftSchool/helper"
)

// LoadRoutes sets up the routes for the server
func (s *Server) LoadRoutes(mux *http.ServeMux) {

	// Initialize shared DB connection once and reuse it for services
	dbConnection, err := s.initializeDBConnection()
	if err != nil {
		helper.SugarObj.Error("Failed to initialize database connection: %v", err)
		return
	}

	// Initialize services with the shared DB connection
	feeService := fees.NewFeeService(dbConnection)

	studentService := student.NewStudentService(dbConnection)

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

// initializeDBConnection initializes and returns the shared database connection
func (s *Server) initializeDBConnection() (*sql.DB, error) {
	// Logic to initialize and return a single DB connection
	// This can be a connection pool or a single DB connection depending on your architecture
	dbConnection, err := helper.OpenPSQLDB(s.Config)
	if err != nil {
		return nil, err
	}
	return dbConnection, nil
}
