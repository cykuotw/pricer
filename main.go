package main

import (
	"log"
	"net/http"
	"pricing-app/api"
	"pricing-app/config"
)

// main is the entry point of the Pricing App.
// It initializes the API server and starts it.
//
// The server address is retrieved from the environment configuration.
// If the server fails to start, the application logs the error and exits.
func main() {
	apiServer := api.NewAPIServer(config.Envs.BackendURL)

	if err := apiServer.Run(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
