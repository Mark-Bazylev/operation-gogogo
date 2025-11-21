package http

import (
	"net/http"

	"operation-gogogo/internal/config"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(applicationConfig config.Config) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	registerHealthRoutes(mux, applicationConfig)
	registerUserRoutes(mux)

	return LoggingMiddleware(mux)
}
