package route

import (
	"go_project_Gin/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.RouterGroup) {
	task := router.Group("tasks")
	{
		task.GET("/", handler.GetAllTasks)
		task.POST("/", handler.CreateTask)
	}
}
