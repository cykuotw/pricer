package pricing

import (
	"fmt"
	"net/http"
	"pricing-app/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func (h *Handler) hGetConfig(c *gin.Context) {
	ticker := strings.ToUpper(c.Param("ticker"))
	config, ok := h.controller.config[ticker]
	if !ok {
		services.WriteJSON(c, http.StatusBadRequest, gin.H{"message": "invalid ticker"})
		return
	}

	services.WriteJSON(c, http.StatusOK, gin.H{"ticker": ticker, "config": config})
}

func (h *Handler) hPostConfig(c *gin.Context) {
	type cfg struct {
		Ticker     string          `json:"ticker"`
		Drift      decimal.Decimal `json:"drift"`
		Volatility decimal.Decimal `json:"volatility"`
	}

	var payload cfg
	if err := services.ParseJSON(c, &payload); err != nil {
		services.WriteJSON(c, http.StatusBadRequest, gin.H{"message": "invalid payload"})
		return
	}

	fmt.Println(h.controller.config)
	_, ok := h.controller.config[payload.Ticker]
	if !ok {
		services.WriteJSON(c, http.StatusBadRequest, gin.H{"message": "invalid ticker"})
		return
	}

	config := h.controller.config[payload.Ticker]
	config.Drift = payload.Drift
	config.Volatility = payload.Volatility
	h.controller.config[payload.Ticker] = config

	services.WriteJSON(c, http.StatusCreated, gin.H{"message": "success"})
}
