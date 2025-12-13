package middleware

import (
	"go_project_Gin/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// basic logger middleware
// func LoggerMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		start := time.Now()

// 		c.Next()

// 		latency := time.Since(start)
// 		statusCode := c.Writer.Status()
// 		clientIP := c.ClientIP()
// 		method := c.Request.Method
// 		path := c.Request.URL.Path

// 		log.Printf("| %3d | %13v | %15s | %s %s",
// 			statusCode,
// 			latency,
// 			clientIP,
// 			method,
// 			path,
// 		)
// 	}
// }

// advanced logger middleware

// var logger *zap.Logger

// func init() {
// 	logger, _ = zap.NewProduction()
// }

// func LoggerMiddleware() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		start := time.Now()
// 		ctx.Next()

// 		logger.Info("HTTP Request",
// 			zap.Int("Status Code", ctx.Writer.Status()),
// 			zap.String("Method", ctx.Request.Method),
// 			zap.String("Path", ctx.Request.URL.Path),
// 			zap.Duration("Latency", time.Since(start)),
// 		)
// 	}
// }

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()

		latency := time.Since(start)
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()
		method := ctx.Request.Method
		path := ctx.Request.URL.Path

		field := []zap.Field{
			zap.Int("status_code", statusCode),
			zap.String("method", method),
			zap.String("path", path),
			zap.Duration("latency", latency),
			zap.String("client_ip", clientIP),
		}

		if statusCode >= 500 {
			utils.Logger.Error("HTTP Request", field...)
			return
		} else if statusCode >= 400 {
			utils.Logger.Warn("HTTP Request", field...)
			return
		}
		utils.Logger.Info("HTTP Request", field...)
	}
}
