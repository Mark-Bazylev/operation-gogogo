package main

//go:generate sh -c "command -v swag >/dev/null 2>&1 && swag init -g ./cmd/api/main.go -o ./docs || echo 'swag not installed; skipping docs generation'"

import (
	"fmt"
	"log"
	"net/http"
	httpRoutes "operation-gogogo/internal/http"
	"os"

	"operation-gogogo/internal/config"

	"operation-gogogo/docs"
)

var buildVersion = "dev" // set by -ldflags at build time

// @title Operation GoGoGo API
// @version 1.0
// @description Learning Go backend step by step.

// @host localhost:8080
// @BasePath /
func main() {
	applicationLogger := log.New(os.Stdout, "", log.LstdFlags)

	applicationConfig := config.LoadConfig(buildVersion)

	// Initialize Swagger metadata at runtime
	docs.SwaggerInfo.Title = "Operation GoGoGo API"
	docs.SwaggerInfo.Version = applicationConfig.Version
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", applicationConfig.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}

	printStartupLine(applicationLogger, applicationConfig)

	router := httpRoutes.NewRouter(applicationConfig)

	serverAddress := fmt.Sprintf(":%d", applicationConfig.Port)

	startError := http.ListenAndServe(serverAddress, router)
	if startError != nil {
		applicationLogger.Fatalf("server failed to start: %v", startError)
	}
}

func printStartupLine(applicationLogger *log.Logger, applicationConfig config.Config) {
	green := "\033[32m"
	reset := "\033[0m"

	coloredPort := fmt.Sprintf("%s%d%s", green, applicationConfig.Port, reset)

	applicationURL := fmt.Sprintf("http://localhost:%d", applicationConfig.Port)

	applicationLogger.Printf(
		"Running on %s (%s)  (env: %s, version: %s)",
		applicationURL,
		coloredPort,
		applicationConfig.Environment,
		applicationConfig.Version,
	)
}
