package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port        int
	Environment string
	Version     string
}

func LoadConfig(buildVersion string) Config {
	return Config{
		Port:        loadPort(),
		Environment: loadEnvironment(),
		Version:     loadVersion(buildVersion),
	}
}

func loadPort() int {
	portString, exists := os.LookupEnv("PORT")

	// Missing PORT → use default
	if !exists || portString == "" {
		return 8080
	}

	portNumber, conversionError := strconv.Atoi(portString)

	// Invalid PORT → warn and use default
	if conversionError != nil {
		log.Printf("Warning: PORT=%q is not a valid number. Using 8080 instead.", portString)
		return 8080
	}

	return portNumber
}

func loadEnvironment() string {
	environmentValue, exists := os.LookupEnv("ENVIRONMENT")
	if !exists || environmentValue == "" {
		return "development"
	}
	return environmentValue
}

func loadVersion(buildVersion string) string {
	if buildVersion == "" {
		return "dev"
	}
	return buildVersion
}
