package route

import (
	"go_project_Gin/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", handler.RegisterHandler)
	router.POST("/login", handler.LoginHandler)

}
