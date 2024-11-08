package main

import (
	"github.com/chinmayvivek/SwiftSchool/config"
	"github.com/chinmayvivek/SwiftSchool/helper"
	"github.com/chinmayvivek/SwiftSchool/server"
)

func main() {
	cfg := config.NewConfig()

	// Initialize and start the server (no need to capture the return value)
	helper.SugarObj.Info("Starting the server...")
	server.Run(cfg)
}
