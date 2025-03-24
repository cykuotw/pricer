package router

import (
	"pricing-app/services/pricing/controller"

	"github.com/gin-gonic/gin"
)

// Handler is responsible for registering routes and delegating requests to the appropriate controller methods.
type Handler struct {
	controller *controller.Contoller
}

// NewHandler creates a new instance of Handler.
//
// Parameters:
// - controller: A pointer to the Contoller instance that contains the business logic.
//
// Returns:
// - A pointer to the initialized Handler instance.
func NewHandler(controller *controller.Contoller) *Handler {
	return &Handler{
		controller: controller,
	}
}

// RegisterRoutes registers all the API routes for the pricing service.
//
// Parameters:
// - router: A Gin router group to which the routes will be added.
//
// Routes:
// - GET `/tickers`: Retrieves a list of all available stock tickers.
// - GET `/config/:ticker`: Fetches the configuration (drift and volatility) for a specific ticker.
// - PUT `/config`: Updates the configuration (drift and volatility) for a specific ticker.
// - GET `/check-open`: Checks if the market is currently open.
// - GET `/prices/:ticker`: Retrieves historical prices for a specific ticker.
// - GET `/stream/server-time`: Streams the current server time to clients using Server-Sent Events (SSE).
// - GET `/stream/update-price/:ticker`: Streams live price updates for a specific ticker using SSE.
func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/tickers", h.hGetTickerList)

	router.GET("/config/:ticker", h.hGetConfig)
	router.PUT("/config", h.hPostConfig)

	router.GET("/check-open", h.hCheckMarketOpen)

	router.GET("/prices/:ticker", h.hGetPrices)

	router.GET("/stream/server-time", h.hStreamServerTime)
	router.GET("/stream/update-price/:ticker", h.hStreamUpdatePrice)
}
