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

// APIServer represents the HTTP server for the Pricing App API.
type APIServer struct {
	addr   string       // Address where the server will listen.
	engine *http.Server // HTTP server instance.
}

// NewAPIServer creates a new instance of APIServer.
// It takes the server address as a parameter.
//
// Example:
//
//	server := NewAPIServer(":8080")
func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

// Run starts the API server.
// It initializes the market configuration, sets up the Gin router with middleware,
// and registers the pricing service routes.
//
// Returns an error if the server fails to start.
func (s *APIServer) Run() error {
	// Load market configuration from a JSON file.
	marketConfig, err := config.LoadMarketConfig("data/initParam.json")
	if err != nil {
		log.Fatal(err)
	}

	// Set Gin mode (debug or release) based on environment configuration.
	gin.SetMode(config.Envs.Mode)

	// Create a new Gin router and apply middleware.
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.CORSMiddleware())

	// Define a sub-router group for API endpoints.
	subRouter := router.Group(config.Envs.APIPath)

	// Create the pricing controller and handler with dependency injection.
	controller := priceController.NewController(marketConfig)
	handler := priceRouter.NewHandler(controller)
	handler.RegisterRoutes(subRouter)

	// Log the server address.
	log.Println("API Server Listening on", s.addr)

	// Configure the HTTP server.
	s.engine = &http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	// Start the server and listen for incoming requests.
	return s.engine.ListenAndServe()
}
