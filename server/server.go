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
	srv.StartServer()
}

// StartServer initializes and starts the HTTP server with proper error handling
func (s *Server) StartServer() {
	// Create a new mux router
	mux := http.NewServeMux()

	// Set up routes
	s.LoadRoutes(mux)

	// Read server timeout from the configuration
	serverTimeout := s.Config.App.ServerTimeout
	readHeaderTimeout := s.Config.App.ReadHeaderTimeout

	// Create the server with timeouts
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", s.Config.App.ServerPort),
		Handler:           mux,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      serverTimeout,
		IdleTimeout:       serverTimeout,
	}

	// Graceful shutdown using channels
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		helper.SugarObj.Info("Server is starting on port %d...", s.Config.App.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			helper.SugarObj.Error("Error starting server: %v", err)
		}
	}()

	// Wait for a termination signal
	<-stop

	// Gracefully shut down the server
	helper.SugarObj.Info("Shutting down the server...")

	// Create a context for graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		helper.SugarObj.Error("Error shutting down the server: %v", err)
	} else {
		helper.SugarObj.Info("Server gracefully stopped.")
	}
}
