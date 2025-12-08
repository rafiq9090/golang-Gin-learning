package route

import (
	"go_project_Gin/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.RouterGroup) {
	public := router.Group("")
	{
		public.GET("/tasks", handler.GetAllTasks)
		public.POST("/tasks", handler.CreateTask)
	}
}
