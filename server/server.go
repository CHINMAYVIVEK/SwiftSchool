package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chinmayvivek/SwiftSchool/config"
	"github.com/chinmayvivek/SwiftSchool/helper"
)

type Server struct {
	Config *config.Config
}

func Run(cfg *config.Config) {
	// Initialize the server with the configuration
	srv := &Server{
		Config: cfg,
	}
	// Start the server
	srv.Start()
}

// Start initializes and starts the HTTP server with proper error handling
func (s *Server) Start() {
	// Set up the mux router and load routes
	mux := http.NewServeMux()
	s.LoadRoutes(mux)

	// Configure the HTTP server with proper timeouts
	server := s.createHTTPServer(mux)

	// Set up graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go s.startServer(server)

	// Wait for termination signal and shut down gracefully
	s.handleShutdown(server, stop)
}

// createHTTPServer creates and configures an HTTP server with timeouts
func (s *Server) createHTTPServer(mux http.Handler) *http.Server {
	return &http.Server{
		Addr:              fmt.Sprintf(":%d", s.Config.App.ServerPort),
		Handler:           mux,
		ReadHeaderTimeout: s.Config.App.ReadHeaderTimeout,
		WriteTimeout:      s.Config.App.ServerTimeout,
		IdleTimeout:       s.Config.App.ServerTimeout,
	}
}

// startServer starts the HTTP server and logs any errors
func (s *Server) startServer(server *http.Server) {
	helper.SugarObj.Info("Server starting on port %d...", s.Config.App.ServerPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		helper.SugarObj.Error("Error starting server: %v", err)
	}
}

// handleShutdown gracefully shuts down the server on receiving a termination signal
func (s *Server) handleShutdown(server *http.Server, stop chan os.Signal) {
	// Wait for termination signal (SIGINT or SIGTERM)
	<-stop

	helper.SugarObj.Info("Shutting down the server...")

	// Create a context with a timeout for graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Perform the shutdown
	if err := server.Shutdown(shutdownCtx); err != nil {
		helper.SugarObj.Error("Error shutting down the server: %v", err)
	} else {
		helper.SugarObj.Info("Server gracefully stopped.")
	}
}
