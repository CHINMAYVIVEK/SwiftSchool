package server

import (
	"swiftschool/app/core"
)

// SetupRoutes initializes all the routes for the server
func (s *Server) SetupRoutes() {
	// Health Check Route
	s.mux.HandleFunc("/api/health", s.handleHealthCheck)

	// Institutes
	institutesService := core.NewService(s.db)
	institutesHandler := core.NewHandler(institutesService)

	// Register Institute
	s.mux.HandleFunc("/api/institutes/register", institutesHandler.CreateInstitute)
}
