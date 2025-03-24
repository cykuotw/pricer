package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config represents the environment configuration for the application.
type Config struct {
	Mode       string // The mode of the application (e.g., "debug" or "release").
	BackendURL string // The URL of the backend service.
	APIPath    string // The base path for API endpoints.

	CORSFrontendOrigin string // The allowed origin for CORS requests from the frontend.
}

// Envs holds the initialized configuration values for the application.
var Envs = initConfig()

// initConfig initializes the application configuration by loading environment variables.
// It uses default values if the environment variables are not set.
func initConfig() Config {
	godotenv.Load()

	return Config{
		Mode:       getEnv("MODE", "debug"),
		BackendURL: getEnv("BACKEND_URL", "localhost:8000"),
		APIPath:    getEnv("API_URL", ""),

		CORSFrontendOrigin: getEnv("CORS_FRONTEND_ORIGIN", "localhost:8050"),
	}
}

// getEnv retrieves the value of an environment variable.
// If the variable is not set, it returns the provided fallback value.
func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
