package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/chinmayvivek/SwiftSchool/config"
	"github.com/chinmayvivek/SwiftSchool/server"
)

func main() {
	// Check if '--build' argument is passed
	if len(os.Args) > 1 && os.Args[1] == "--build" {
		fmt.Println("Building the project...")
		runBuild() // Call the build function if '--build' is passed
		return
	}

	cfg := config.NewConfig()

	fmt.Println("Configuration loaded successfully.")
	fmt.Println("Starting the server...")

	server.Run(cfg)
}

// Helper function to execute a command and handle errors
func executeCommand(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command and check for errors
	if err := cmd.Run(); err != nil {
		log.Fatalf("Command failed: %v\n", err)
	}
}

func runBuild() {
	// --- Go Build for macOS ---
	fmt.Println("Building Go project for macOS...")

	// Set environment variables for macOS build
	cmd := exec.Command("go", "build", "-o", "build/swift_school-macos", "-ldflags", "-s -w", "./")
	cmd.Env = append(os.Environ(), "GOOS=darwin", "GOARCH=amd64")
	executeCommand(cmd)

	// --- Go Build for Windows ---
	fmt.Println("Building Go project for Windows...")

	// Set environment variables for Windows build
	cmd = exec.Command("go", "build", "-o", "build/swift_school-windows.exe", "-ldflags", "-s -w", "./")
	cmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
	executeCommand(cmd)

	// --- Flutter Build for macOS ---
	// Change directory to the Flutter project directory (flutter_project)
	fmt.Println("Building Flutter project for macOS...")
	cmd = exec.Command("flutter", "build", "macos", "--release", "--no-dart-define", "-C", "swift_school")
	executeCommand(cmd)

	// --- Flutter Build for Windows ---
	fmt.Println("Building Flutter project for Windows...")
	cmd = exec.Command("flutter", "build", "windows", "--release", "--no-dart-define", "-C", "swift_school")
	executeCommand(cmd)

	// --- Finished ---
	fmt.Println("Build completed for macOS and Windows!")
}
