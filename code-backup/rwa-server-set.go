package main

/*

package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

func initLogger() *slog.Logger {

	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	return logger
}

// Middleware function
func handleMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware: Request received for", r.URL.Path)

		// Generate a new UUID as Request-ID
		requestID := uuid.New().String()

		// Set headers

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Request-ID", requestID)
		w.Header().Set("name", "santo")

		// Add requestID to request header (for handler to access)
		r.Header.Set("X-Request-ID", requestID)
		r.Header.Set("name", "Santo")

		slog.Info("Incoming request",
			"method", r.Method,
			"path", r.URL.Path,
			"request_id", requestID,
			"remote_adr", r.RemoteAddr,
		)

		// Log the Request ID (for debugging/logging purposes)
		fmt.Println("Request ID:", requestID)

		// Call next handler
		next.ServeHTTP(w, r)

		duration := time.Since(time.Now())
		slog.Info("Complete request",
			"request_id", requestID,
			"duration", duration,
		)
	}
}

// Home handler (returns JSON)
func homeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Welcome to the web server!",
	}
	json.NewEncoder(w).Encode(response)
}

// About handler (returns JSON with request ID)
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-ID")
	name := r.Header.Get("name")

	response := map[string]string{
		"message":    "This is the About page.",
		"request_id": requestID,
		"names ":     name,
	}
	json.NewEncoder(w).Encode(response)
}

func main() {

	initLogger()
	// Route setup with middleware
	http.HandleFunc("/", handleMiddleware(homeHandler))
	http.HandleFunc("/about", handleMiddleware(aboutHandler))

	// Start server
	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}



*/