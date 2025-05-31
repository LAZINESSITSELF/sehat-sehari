package middlewares

import (
    "log"
    "time"

    "github.com/gin-gonic/gin"
)

// LoggerMiddleware mencatat method, path, status, dan latensi.
func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        startTime := time.Now()
        c.Next()
        latency := time.Since(startTime)
        log.Printf("Method: %s, Path: %s, Status: %d, Latency: %s",
            c.Request.Method, c.Request.RequestURI, c.Writer.Status(), latency)
    }
}