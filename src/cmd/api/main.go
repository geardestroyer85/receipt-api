package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	fmt.Println("Starting server on port 8080...")
	if err := engine.Run(":8080"); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
