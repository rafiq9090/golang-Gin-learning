package route

import (
	"go_project_Gin/internal/handler"
	"go_project_Gin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.RouterGroup) {
	task := router.Group("tasks")
	{

		task.GET("/", handler.GetAllTasks)
		protected := task.Group("")
		protected.Use(middleware.JWTAuthMiddleware())
		protected.POST("/", handler.CreateTask)
		protected.PUT("/:id", handler.UpdateTask)
		protected.DELETE("/:id", handler.DeleteTask)

	}
}
