package api

import (
	"log"
	"net/http"
	"pricing-app/config"
	"pricing-app/services/pricing"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	addr string

	engine *http.Server
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	gin.SetMode("debug")

	router := gin.New()
	router.Use(gin.Logger())

	subRouter := router.Group(config.Envs.APIPath)

	homeHandler := pricing.NewHandler()
	homeHandler.RegisterRoutes(subRouter)

	log.Println("API Server Listening on", s.addr)

	s.engine = &http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	return s.engine.ListenAndServe()
}
