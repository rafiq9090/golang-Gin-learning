package main

import (
	"go_project_Gin/internal/config"
	"go_project_Gin/internal/database"
	"go_project_Gin/internal/middleware"
	"go_project_Gin/internal/route"
	"go_project_Gin/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	config.Load()
	utils.InitLogger()
	defer utils.Logger.Sync()
	database.Connect()

	r := gin.New()

	api := r.Group("/api")
	r.Use(middleware.LoggerMiddleware())
	route.SetupRoutes(api)
	utils.Logger.Info("Server running on port", zap.String("port", config.App.Port))
	if err := r.Run(":" + config.App.Port); err != nil {
		utils.Logger.Fatal("Server failed to start", zap.Error(err))
	}
}
