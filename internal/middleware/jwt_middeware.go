package middleware

import (
	"go_project_Gin/internal/config"
	"go_project_Gin/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	UserId uint `json:"user_id"`
	jwt.RegisteredClaims
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.JSONError(c, "Authorization header required", http.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.JSONError(c, "Invalid Authorization header", http.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		claims := JWTClaim{}
		tokenString := parts[1]
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.App.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			utils.JSONError(c, "Invalid or expired token", http.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserId)
		c.Next()
	}
}
