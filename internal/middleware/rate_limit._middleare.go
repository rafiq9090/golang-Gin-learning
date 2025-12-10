package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	visitor = make(map[string]*rate.Limiter)
	mu      sync.Mutex
)

func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitor[ip]

	if !exists {
		limiter = rate.NewLimiter(rate.Every(time.Second/2), 5)
		visitor[ip] = limiter
	}
	return limiter
}

func RateLimitMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		limiter := getLimiter(ip)

		if !limiter.Allow() {
			ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}

}
