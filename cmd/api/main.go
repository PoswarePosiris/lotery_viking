package main

import (
	"fmt"
	"log"
	"lotery_viking/internal/database"
	"lotery_viking/internal/server"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [migrate|drop|seed|serve]")
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
	default:
		fmt.Println("Usage: go run main.go [migrate|drop|seed|serve]")
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
