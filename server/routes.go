package server

import (
	"swiftschool/app/institutes"
)

// SetupRoutes initializes all the routes for the server
func (s *Server) SetupRoutes() {
	// Health Check Route
	s.mux.HandleFunc("/api/health", s.handleHealthCheck)

	// Institutes
	institutesService := institutes.NewService(s.db)
	institutesHandler := institutes.NewHandler(institutesService)

	// Register Institute
	s.mux.HandleFunc("/api/institutes/register", institutesHandler.InstitutesRegistration)
}
