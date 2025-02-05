package main

import (
	"log"
	"receipt-api/src/internal/config"
	"receipt-api/src/internal/infrastructure/rest/handlers"
	"receipt-api/src/internal/infrastructure/rest/router"
	"receipt-api/src/internal/infrastructure/rest/server"
)

func main() {

	config := config.LoadConfig()

	receiptHandler := handlers.NewReceiptHandler()
	router := router.NewRouter()

	server := server.NewServer(router)
	server.SetupRoutes()

	port := config.AppPort
	if err := server.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Printf("Server started on port %s\n", port)

}
