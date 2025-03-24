package router

import (
	"net/http"
	"pricing-app/services"

	"github.com/gin-gonic/gin"
)

// hGetTickerList handles the HTTP GET request to retrieve a list of all available stock tickers.
//
// This endpoint responds with a JSON object containing the list of tickers.
//
// Example response:
//
//	{
//	    "tickers": ["AAPL", "GOOG", "MSFT"]
//	}
//
// Parameters:
// - c: The Gin context for the HTTP request.
func (h *Handler) hGetTickerList(c *gin.Context) {
	tickers := h.controller.GetTickers()

	services.WriteJSON(c, http.StatusOK, gin.H{"tickers": tickers})
}
