package http

import (
	"net/http"

	"operation-gogogo/internal/config"
)

func NewRouter(applicationConfig config.Config) http.Handler {
	mux := http.NewServeMux()

	healthCheckHandler := &HealthCheckHandler{
		ApplicationConfig: applicationConfig,
	}

	mux.Handle("/health", healthCheckHandler)

	return mux
}
