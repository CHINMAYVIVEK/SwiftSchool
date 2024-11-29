package server

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chinmayvivek/SwiftSchool/config"
	"github.com/chinmayvivek/SwiftSchool/helper"
	"github.com/rs/cors"
)

type Server struct {
	Config *config.Config
	DB     *sql.DB
}

func Run(cfg *config.Config) {
	db, err := cfg.NewPSQLDBConnection()
	if err != nil {
		helper.SugarObj.Error("Failed to connect to database: ", err)
		return
	}
	helper.SugarObj.Info("Connected to database.")

	server := &Server{Config: cfg, DB: db}
	server.Start()
}

func (s *Server) Start() {
	mux := http.NewServeMux()
	s.LoadRoutes(mux)

	// Configure CORS with specific options for Flutter desktop
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:           true,
	}

	// Create a new CORS handler with our options
	handler := cors.New(corsOptions).Handler(mux)

	// Explicitly use IPv4
	addr := fmt.Sprintf("127.0.0.1:%d", s.Config.App.ServerPort)
	server := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: s.Config.App.ReadHeaderTimeout,
		WriteTimeout:      s.Config.App.ServerTimeout,
		IdleTimeout:       s.Config.App.ServerTimeout,
	}

	// Print server info
	helper.SugarObj.Info("Server configuration:")
	helper.SugarObj.Info("- Listening on: %s", addr)
	helper.SugarObj.Info("- CORS enabled with all origins (*)")
	helper.SugarObj.Info("- Debug mode: enabled")
	helper.SugarObj.Info("- IPv4 mode: enabled")
	helper.SugarObj.Info("- Available endpoints:")
	helper.SugarObj.Info("  * /api/health")
	helper.SugarObj.Info("  * /api/data")
	helper.SugarObj.Info("  * /api/fees/fee-structure-by-class")
	helper.SugarObj.Info("  * /api/student/student-registration")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		helper.SugarObj.Info("Starting server on %s...", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			helper.SugarObj.Error("Server error: %v", err)
		}
	}()

	<-stop
	helper.SugarObj.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		helper.SugarObj.Error("Shutdown error: %v", err)
	} else {
		helper.SugarObj.Info("Server stopped gracefully")
	}
}

