package main

import (
	"log"
	"net/http"
	"pricing-app/api"
	"pricing-app/config"
)

func main() {
	apiServer := api.NewAPIServer(config.Envs.BackendURL)

	if err := apiServer.Run(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
