package server

import (
	"swiftschool/app/core"
)

// SetupRoutes initializes all the routes for the server
func (s *Server) SetupRoutes() {
	// Health Check Route
	s.mux.HandleFunc("/api/health", s.handleHealthCheck)

	// Institutes
	coreService := core.NewService(s.db)
	coreHandler := core.NewHandler(coreService)

	// Register Institute
	s.mux.HandleFunc("/api/institutes/register", coreHandler.CreateInstitute)
}
