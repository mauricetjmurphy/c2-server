package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Calculate resolution time
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// Get status code
		statusCode := c.Writer.Status()

		// Log details
		log.Printf("%s %s %d %s",
			c.Request.Method,
			c.Request.URL.Path,
			statusCode,
			latency,
		)
	}
}
