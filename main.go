package main

import (
	"log"

	"github.com/chinmayvivek/SwiftSchool/config"
	"github.com/chinmayvivek/SwiftSchool/server"
)

func main() {
	cfg := config.NewConfig()
	// server.Run(cfg)

	// Initialize and start the server (no need to capture the return value)
	log.Println("Starting the server...")
	server.Run(cfg)
}
