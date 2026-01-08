package route

import (
	"go_project_Gin/internal/handler"
	"go_project_Gin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(router *gin.RouterGroup) {
	post := router.Group("posts")
	post.Use(middleware.JWTAuthMiddleware())
	{
		post.POST("/", handler.CreatePost)
		post.GET("/my", handler.GetPostsByUserId)
		post.GET("/", handler.GetAllPost)
	}
}
