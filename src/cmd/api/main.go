package main

import (
	"log"
	"receipt-api/src/internal/config"
	"receipt-api/src/internal/infrastructure/rest/server"
)

func main() {

	config := config.LoadConfig()

	server := server.NewServer()
	server.SetupRoutes()

	port := config.AppPort
	if err := server.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Printf("Server started on port %s\n", port)

}
