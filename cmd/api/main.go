package main

import (
	"fmt"
	"lotery_viking/internal/server"
	"os"
)

func main() {

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
