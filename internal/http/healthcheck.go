package http

import (
	"encoding/json"
	"net/http"

	"operation-gogogo/internal/config"
)

type HealthCheckHandler struct {
	ApplicationConfig config.Config
}

type HealthCheckResponse struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func (handler *HealthCheckHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	response := HealthCheckResponse{
		Status:      "ok",
		Environment: handler.ApplicationConfig.Environment,
		Version:     handler.ApplicationConfig.Version,
	}

	responseWriter.Header().Set("Content-Type", "application/json")

	jsonEncoder := json.NewEncoder(responseWriter)
	jsonEncoder.SetIndent("", "  ") //Purely cosmetic addon to make json have spaces.
	err := jsonEncoder.Encode(response)
	if err != nil {
		http.Error(responseWriter, "failed to encode health response", http.StatusInternalServerError)
		return
	}
}
