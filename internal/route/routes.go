package route

import (
	"go_project_Gin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(middleware.RateLimitMiddleware())
	r.Use(middleware.CORSMiddleware())

	api := r.Group("/api")
	SetupTaskRoutes(api)
}
