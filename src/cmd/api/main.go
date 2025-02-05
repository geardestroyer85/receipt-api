package main

import (
	"log"
	"receipt-api/src/internal/application/services"
	"receipt-api/src/internal/config"
	"receipt-api/src/internal/infrastructure/rest/handlers"
	"receipt-api/src/internal/infrastructure/rest/router"
	"receipt-api/src/internal/infrastructure/rest/server"
	"receipt-api/src/internal/infrastructure/storage/memory"
)

func main() {

	config := config.LoadConfig()

	receiptRepo := memory.NewMemoryReceiptRepository()
	receiptService := services.NewReceiptService(receiptRepo)
	receiptHandler := handlers.NewReceiptHandler(receiptService)

	router := router.NewRouter(receiptHandler)
	server := server.NewServer(router)
	server.SetupRoutes()

	port := config.AppPort
	if err := server.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Printf("Server started on port %s\n", port)

}
