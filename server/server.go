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
)

type Server struct {
	Config *config.Config
	DB     *sql.DB
}

func Run(cfg *config.Config) {
	db, err := config.NewConfig().NewPSQLDBConnection()
	if err != nil {
		helper.SugarObj.Error("Failed to connect to database: %v", err)
		return
	}
	helper.SugarObj.Info("Connected to database.")

	server := &Server{Config: cfg, DB: db}
	server.Start()
}

func (s *Server) Start() {
	mux := http.NewServeMux()
	s.LoadRoutes(mux)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", s.Config.App.ServerPort),
		Handler:           mux,
		ReadHeaderTimeout: s.Config.App.ReadHeaderTimeout,
		WriteTimeout:      s.Config.App.ServerTimeout,
		IdleTimeout:       s.Config.App.ServerTimeout,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go s.runServer(server)
	s.shutdown(server, stop)
}

func (s *Server) runServer(server *http.Server) {
	helper.SugarObj.Info("Starting server on port %d...", s.Config.App.ServerPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		helper.SugarObj.Error("Server error: %v", err)
	}
}

func (s *Server) shutdown(server *http.Server, stop chan os.Signal) {
	<-stop
	helper.SugarObj.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		helper.SugarObj.Error("Shutdown error: %v", err)
	} else {
		helper.SugarObj.Info("Server stopped.")
	}
}
