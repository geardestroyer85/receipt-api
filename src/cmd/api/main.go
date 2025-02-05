package main

import (
	"fmt"
	"receipt-api/src/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {

	config := config.LoadConfig()

	engine := gin.New()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	fmt.Println("Starting server on port 8080...")
	if err := engine.Run(fmt.Sprintf(":%s", config.AppPort)); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

}
