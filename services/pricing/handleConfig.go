package pricing

import (
	"net/http"
	"pricing-app/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

var DRIFT_MULTIPLE = decimal.NewFromInt32(10000)
var VOLATILITY_MULTIPLE = decimal.NewFromInt32(100)

func (h *Handler) hGetConfig(c *gin.Context) {
	ticker := strings.ToUpper(c.Param("ticker"))
	config, ok := h.controller.config[ticker]
	if !ok {
		services.WriteJSON(c, http.StatusBadRequest, gin.H{"message": "invalid ticker"})
		return
	}

	shiftedConfig := config

	shiftedConfig.Drift = shiftedConfig.Drift.Mul(DRIFT_MULTIPLE)
	shiftedConfig.Volatility = shiftedConfig.Volatility.Mul(VOLATILITY_MULTIPLE)

	services.WriteJSON(c, http.StatusOK, gin.H{"ticker": ticker, "config": shiftedConfig})
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

	_, ok := h.controller.config[payload.Ticker]
	if !ok {
		services.WriteJSON(c, http.StatusBadRequest, gin.H{"message": "invalid ticker"})
		return
	}

	config := h.controller.config[payload.Ticker]
	config.Drift = payload.Drift.Div(DRIFT_MULTIPLE)
	config.Volatility = payload.Volatility.Div(VOLATILITY_MULTIPLE)
	h.controller.config[payload.Ticker] = config

	services.WriteJSON(c, http.StatusCreated, gin.H{"message": "success"})
}
