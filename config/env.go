package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Mode       string
	BackendURL string
	APIPath    string

	CORSFrontendOrigin string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Mode:       getEnv("MODE", "debug"),
		BackendURL: getEnv("BACKEND_URL", "localhost:8000"),
		APIPath:    getEnv("API_URL", ""),

		CORSFrontendOrigin: getEnv("CORS_FRONTEND_ORIGIN", "localhost:8050"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
