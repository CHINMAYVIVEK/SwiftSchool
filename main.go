package main

import (
	"github.com/chinmayvivek/SwiftSchool/config"
	"github.com/chinmayvivek/SwiftSchool/helper"
	"github.com/chinmayvivek/SwiftSchool/server"
)

func main() {
	cfg := config.NewConfig()
	helper.SugarObj.Info("Starting the server...")

	server.Run(cfg)
}
