package http

import (
	"net/http"

	"operation-gogogo/internal/config"
)

func registerHealthRoutes(mux *http.ServeMux, applicationConfig config.Config) {
	healthCheckHandler := &HealthCheckHandler{
		ApplicationConfig: applicationConfig,
	}
	mux.Handle("/health", healthCheckHandler)
}
