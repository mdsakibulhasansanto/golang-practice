package main 

/*

package main

import (
	"log/slog"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Logger initialization
func initLogger() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}

// Gin middleware function
func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()

		// Set response headers
		c.Header("Content-Type", "application/json")
		c.Header("X-Request-ID", requestID)
		c.Header("name", "santo")

		// Set headers to request context (if needed)
		c.Request.Header.Set("X-Request-ID", requestID)
		c.Request.Header.Set("name", "Santo")

		slog.Info("Incoming request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"request_id", requestID,
			"remote_addr", c.ClientIP(),
		)

		start := time.Now()

		// Process request
		c.Next()

		duration := time.Since(start)

		slog.Info("Complete request",
			"request_id", requestID,
			"duration", duration,
		)
	}
}

// Handler: Home
func homeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the web server!",
	})
}

// Handler: About
func aboutHandler(c *gin.Context) {
	requestID := c.GetHeader("X-Request-ID")
	name := c.GetHeader("name")

	list := []map[string]interface{}{
		{
			"message":    "This is the About page.",
			"request_id": requestID,
			"name":       name,
		},
		{
			"message":    "Another message",
			"request_id": "xyz456",
			"name":       "Nurul",
		},
	}


	c.JSON(200, gin.H{
		"message":    "This is the About page.",
		"request_id": requestID,
		"name":       name,
	})
}

// Main function
func main() {
	initLogger()

	r := gin.Default()
	r.Use(middleware())

	r.GET("/", homeHandler)
	r.GET("/about", aboutHandler)

	slog.Info("Server is running at http://localhost:8080")
	err := r.Run(":8080")
	if err != nil {
		slog.Error("Error starting server", "error", err)
	}
}



*/