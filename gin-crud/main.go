package main

import (
	"go-learn/gin-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupBookRoutes(router)

	// Add a simple health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Run the server on port 8080
	router.Run(":8080")
}