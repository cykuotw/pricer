package router

import (
	"net/http"
	"pricing-app/services"
	"time"

	"github.com/gin-gonic/gin"
)

// hCheckMarketOpen handles the HTTP request to check if the market is currently open.
//
// This endpoint responds with a JSON object indicating whether the market is open.
//
// Example response:
//
//	{
//	    "isopen": true
//	}
//
// Parameters:
// - c: The Gin context for the HTTP request.
func (h *Handler) hCheckMarketOpen(c *gin.Context) {
	isOpen := h.controller.CheckMarketOpen(time.Now())

	services.WriteJSON(c, http.StatusOK, gin.H{"isopen": isOpen})
}
