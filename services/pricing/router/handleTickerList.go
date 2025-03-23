package router

import (
	"net/http"
	"pricing-app/services"

	"github.com/gin-gonic/gin"
)

func (h *Handler) hGetTickerList(c *gin.Context) {
	tickers := h.controller.GetTickers()

	services.WriteJSON(c, http.StatusOK, gin.H{"tickers": tickers})
}
