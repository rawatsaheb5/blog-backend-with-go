package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinLogger returns a gin middleware for logging HTTP requests
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get logger fields
		fields := []zap.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
			zap.String("latency_human", latency.String()),
		}

		// Add error if any
		if len(c.Errors) > 0 {
			fields = append(fields, zap.Strings("errors", c.Errors.Errors()))
		}

		// Log based on status code
		if c.Writer.Status() >= 500 {
			GetLogger().Error("HTTP Request", fields...)
		} else if c.Writer.Status() >= 400 {
			GetLogger().Warn("HTTP Request", fields...)
		} else {
			GetLogger().Info("HTTP Request", fields...)
		}
	}
}

// GinRecovery returns a gin middleware for recovering from panics
func GinRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		GetLogger().Error("Panic recovered",
			zap.Any("error", recovered),
			zap.String("path", c.Request.URL.Path),
			zap.String("method", c.Request.Method),
			zap.String("ip", c.ClientIP()),
		)
		c.AbortWithStatusJSON(500, gin.H{
			"error": "Internal server error",
		})
	})
}
