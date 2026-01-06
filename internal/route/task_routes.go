package route

import (
	"go_project_Gin/internal/handler"
	"go_project_Gin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.RouterGroup) {
	task := router.Group("tasks")
	task.Use(middleware.JWTAuthMiddleware())
	{
		task.GET("/", handler.GetAllTasks)
		task.POST("/", handler.CreateTask)
		task.PUT("/:id", handler.UpdateTask)
		task.DELETE("/:id", handler.DeleteTask)
	}
}
