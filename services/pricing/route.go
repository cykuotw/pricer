package pricing

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	controller *Contoller
}

func NewHandler(controller *Contoller) *Handler {
	return &Handler{
		controller: controller,
	}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/tickers", h.hGetTickerList)

	router.GET("/config/:ticker", h.hGetConfig)
	router.PUT("config", h.hPostConfig)

	router.GET("/stream/server-time", h.hStreamServerTime)
}
