package middlewares

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs request and response details
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s - %s", r.Method, r.URL.Path, time.Since(start))
	})
}
