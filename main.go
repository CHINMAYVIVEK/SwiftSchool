package main

import (
	"fmt"

	"github.com/chinmayvivek/SwiftSchool/config"
	"github.com/chinmayvivek/SwiftSchool/server"
)

func main() {
	cfg := config.NewConfig()

	fmt.Println("Configuration loaded successfully.")
	fmt.Println("Starting the server...")

	server.Run(cfg)
}
