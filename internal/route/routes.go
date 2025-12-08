package route

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(api *gin.RouterGroup) {
	SetupTaskRoutes(api)
}
