package main

// import (
// 	"embed"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"os/exec"
// 	"path/filepath"
// )

// //go:embed flutter_app/build/*
// var flutterAssets embed.FS

// // BuildFlutter builds the Flutter app for a specific platform.
// func BuildFlutter(target string) {
// 	// Set the environment variable for the build process
// 	flutterEnv := os.Getenv("FLUTTER_APP_PATH")
// 	if flutterEnv == "" {
// 		flutterEnv = "./flutter_app" // Default path for Flutter app
// 	}

// 	// Navigate to the Flutter project directory
// 	err := os.Chdir(flutterEnv)
// 	if err != nil {
// 		log.Fatalf("Error navigating to Flutter project directory: %v", err)
// 	}

// 	// Run the Flutter build based on the target platform
// 	cmd := exec.Command("flutter", "build", target)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	err = cmd.Run()
// 	if err != nil {
// 		log.Fatalf("Error building Flutter app: %v", err)
// 	}

// 	log.Printf("Flutter app built successfully for %s.", target)
// }

// // BuildGo builds the Go backend application
// func BuildGo() {
// 	// Check if environment variables are set for the Go build path
// 	goEnv := os.Getenv("GO_APP_PATH")
// 	if goEnv == "" {
// 		goEnv = "./" // Default path to Go app
// 	}

// 	// Build the Go application
// 	cmd := exec.Command("go", "build", "-o", "final_executable")
// 	cmd.Dir = goEnv
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatalf("Error building Go backend: %v", err)
// 	}

// 	log.Println("Go backend built successfully.")
// }

// // CopyFlutterAssets copies the built Flutter app to the Go backend
// func CopyFlutterAssets() {
// 	// Create a directory to store the Flutter build files
// 	buildDir := "./final_build/flutter_assets"
// 	err := os.MkdirAll(buildDir, os.ModePerm)
// 	if err != nil {
// 		log.Fatalf("Error creating Flutter assets directory: %v", err)
// 	}

// 	// Copy the embedded Flutter files to the Go backend build directory
// 	files, err := flutterAssets.ReadDir("flutter_app/build")
// 	if err != nil {
// 		log.Fatalf("Error reading embedded Flutter assets: %v", err)
// 	}

// 	for _, file := range files {
// 		if !file.IsDir() {
// 			data, err := flutterAssets.ReadFile("flutter_app/build/" + file.Name())
// 			if err != nil {
// 				log.Fatalf("Error reading embedded file %s: %v", file.Name(), err)
// 			}

// 			err = ioutil.WriteFile(filepath.Join(buildDir, file.Name()), data, 0644)
// 			if err != nil {
// 				log.Fatalf("Error writing file %s: %v", file.Name(), err)
// 			}
// 		}
// 	}

// 	log.Println("Flutter assets copied to final build.")
// }

// // BuildFinalExecutable creates the final executable that combines Flutter and Go.
// func BuildFinalExecutable() {
// 	// This function assumes that Go backend and Flutter assets are already in place
// 	log.Println("Building final executable...")

// 	// Build Go binary with embedded Flutter assets
// 	cmd := exec.Command("go", "build", "-o", "final_executable")
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatalf("Error creating final executable: %v", err)
// 	}

// 	log.Println("Final executable created successfully.")
// }

// func build() {
// 	// Parse command line arguments or environment variables for platform choice
// 	if len(os.Args) < 2 {
// 		log.Fatal("Please specify a target platform (e.g., windows, linux, macos).")
// 	}
// 	targetPlatform := os.Args[1]

// 	// Set environment variables for the build process (optional, based on platform)
// 	// You can use a build flag or environment variable to set these

// 	// Build Flutter for the target platform
// 	BuildFlutter(targetPlatform)

// 	// Build the Go backend
// 	BuildGo()

// 	// Copy Flutter assets into Go backend
// 	CopyFlutterAssets()

// 	// Build the final executable
// 	BuildFinalExecutable()

// 	log.Println("Build process completed successfully!")
// }
