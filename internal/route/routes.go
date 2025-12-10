package route

import (
	"go_project_Gin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(api *gin.RouterGroup) {
	api.Use(middleware.RateLimitMiddleware())
	api.Use(middleware.CORSMiddleware())
	SetupTaskRoutes(api)
}
