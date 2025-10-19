package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var rateLimiter = rate.NewLimiter(10, 20)

func RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if rateLimiter.Allow() {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			return
		}
	}
}
