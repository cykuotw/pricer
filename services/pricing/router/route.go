package router

import (
	"pricing-app/services/pricing/controller"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	controller *controller.Contoller
}

func NewHandler(controller *controller.Contoller) *Handler {
	return &Handler{
		controller: controller,
	}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/tickers", h.hGetTickerList)

	router.GET("/config/:ticker", h.hGetConfig)
	router.PUT("config", h.hPostConfig)

	router.GET("/prices/:ticker", h.hGetPrices)

	router.GET("/stream/server-time", h.hStreamServerTime)
	router.GET("/stream/update-price", h.hStreamUpdatePrice)
}
