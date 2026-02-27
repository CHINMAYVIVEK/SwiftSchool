package server

import (
	_ "swiftschool/swagger" // Import generated swagger docs

	httpSwagger "github.com/swaggo/http-swagger"
)

// SetupRoutes initializes all application routes
func (s *Server) SetupRoutes() {
	s.registerPageRoutes()
	s.registerAPIRoutes()
	s.registerSwaggerRoute()
}

// registerSwaggerRoute sets up the Swagger UI endpoint
func (s *Server) registerSwaggerRoute() {
	s.mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
}
