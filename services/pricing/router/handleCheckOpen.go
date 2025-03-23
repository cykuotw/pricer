package router

import (
	"net/http"
	"pricing-app/services"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) hCheckMarketOpen(c *gin.Context) {
	isOpen := h.controller.CheckMarketOpen(time.Now())

	services.WriteJSON(c, http.StatusOK, gin.H{"isopen": isOpen})
}
