package main

import (
	"fmt"
	"log"
	"lotery_viking/internal/database"
	"lotery_viking/internal/server"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [migrate|drop|seed|serve|init]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "migrate":
		err := database.Migrate()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Database migration completed successfully")
	case "drop":
		err := database.Drop()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Database dropped successfully")
	case "seed":
		err := database.Seed()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Database seeding completed successfully")
	case "serve":
		startServer()
	case "init":
		err := initialize()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Project initialization completed successfully")
	default:
		fmt.Println("Usage: go run main.go [migrate|drop|seed|serve|init]")
		os.Exit(1)
	}

}

func startServer() {
	server := server.NewServer()

	// Get the host and port
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost" // Default to localhost if HOST is not set
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if PORT is not set
	}

	// Print the full URL before starting the server
	fmt.Printf("Server starting at http://%s:%s\n", host, port)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

func initialize() error {
	// Get the directory where the binary is located
	execDir, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable directory: %w", err)
	}
	baseDir := filepath.Dir(execDir)

	envFile := filepath.Join(baseDir, ".env")
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		// Create the .env file. With this
		content := []byte("HTTP=http # http or https\nHOST=localhost\nPORT=8080\nAPP_ENV=local\nAPI_KEY=\n\nGIN_MODE=release # debug, test, release\n\nDB_HOST=localhost\nDB_PORT=\nDB_DATABASE=\nDB_USERNAME=\nDB_PASSWORD=\nDB_ROOT_PASSWORD=\n")
		err := os.WriteFile(envFile, content, 0644) // Permissions: rw-r--r--
		if err != nil {
			return fmt.Errorf("failed to create .env file: %w", err)
		}
		fmt.Println(".env file created successfully")
	} else {
		fmt.Println(".env file already exists, skipping creation")
	}

	// Create the images folder
	imagesDir := filepath.Join(baseDir, "kiosk_images")
	err = os.MkdirAll(imagesDir, 0755) // Permissions: rwxr-xr-x
	if err != nil {
		return fmt.Errorf("failed to create kiosk_images directory: %w", err)
	}

	return nil
}
