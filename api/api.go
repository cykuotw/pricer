package api

import (
	"log"
	"net/http"
	"pricing-app/config"
	"pricing-app/services/middleware"
	priceController "pricing-app/services/pricing/controller"
	priceRouter "pricing-app/services/pricing/router"

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
	// load market config
	marketConfig, err := config.LoadMarketConfig("data/initParam.json")
	if err != nil {
		log.Fatal(err)
	}

	// set debug/release mode
	gin.SetMode(config.Envs.Mode)

	// create router, setup middle, define group
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.CORSMiddleware())

	subRouter := router.Group(config.Envs.APIPath)

	// create controller & handler with dependency injection
	controller := priceController.NewController(marketConfig)
	handler := priceRouter.NewHandler(controller)
	handler.RegisterRoutes(subRouter)

	log.Println("API Server Listening on", s.addr)

	s.engine = &http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	return s.engine.ListenAndServe()
}
