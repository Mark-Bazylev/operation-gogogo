package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"operation-gogogo/internal/config"
	httpRoutes "operation-gogogo/internal/http"
)

var buildVersion = "dev" // set by -ldflags at build time

func main() {
	applicationLogger := log.New(os.Stdout, "", log.LstdFlags)

	applicationConfig := config.LoadConfig(buildVersion)

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
