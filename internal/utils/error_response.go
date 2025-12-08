package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error   string            `json:"error"`
	Details map[string]string `json:"details"`
}

func JSONError(c *gin.Context, message string, statusCode int, details map[string]string) {
	response := ErrorResponse{
		Error:   message,
		Details: details,
	}
	c.JSON(statusCode, response)
}

func JSONSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
