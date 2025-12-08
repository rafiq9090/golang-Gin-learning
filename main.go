package main

import (
	"go_project_Gin/internal/config"
	"go_project_Gin/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.Load()

	// Connect to database
	database.Connect()

	// Initialize Gin router
	r := gin.Default()

	// Simple health check route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Run the server
	r.Run(":" + config.App.Port)
}
