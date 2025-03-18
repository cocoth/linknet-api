package utils

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func GinCustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()

		log.Printf("[GIN] %s | %3d | %13v | %-15s | %-7s %#v\n",
			time.Now().Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
