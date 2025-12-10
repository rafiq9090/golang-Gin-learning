package main

import (
	"go_project_Gin/internal/config"
	"go_project_Gin/internal/database"
	"go_project_Gin/internal/route"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	database.Connect()

	r := gin.Default()

	route.SetupRoutes(r)
	log.Println("Server running on port", config.App.Port)
	log.Fatal(r.Run(":" + config.App.Port))
}
