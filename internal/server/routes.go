package server

import (
	"github.com/chinmayvivek/swiftschool/app/student"
)

// SetupRoutes initializes all the routes for the server
func (s *Server) SetupRoutes() {
	// Health Check Route
	s.mux.HandleFunc("/api/health", s.handleHealthCheck)

	// Student Routes
	studentRepo := student.NewRepository(s.db)
	studentHandler := student.NewHandler(studentRepo)

	s.mux.HandleFunc("/api/students", studentHandler.CreateStudent)

}
