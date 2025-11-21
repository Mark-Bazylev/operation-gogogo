package http

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		startTime := time.Now().UTC()

		log.Printf("[REQUEST] %s %s", request.Method, request.URL.Path)

		next.ServeHTTP(responseWriter, request)

		elapsed := time.Since(startTime)
		log.Printf("[RESPONSE] %s %s completed in %s", request.Method, request.URL.Path, elapsed)
	})
}
