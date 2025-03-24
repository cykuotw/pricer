package router

import (
	"net/http"
	"pricing-app/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

var DRIFT_MULTIPLE = decimal.NewFromInt32(10000)    // Multiplier for drift values to scale them for API responses.
var VOLATILITY_MULTIPLE = decimal.NewFromInt32(100) // Multiplier for volatility values to scale them for API responses.

// hGetConfig handles the HTTP GET request to retrieve the configuration of a specific stock ticker.
//
// This endpoint responds with the stock's drift and volatility values, scaled for readability.
//
// Example response:
//
//	{
//	    "ticker": "AAPL",
//	    "config": {
//	        "drift": 0.05,
//	        "volatility": 0.02
//	    }
//	}
//
// Parameters:
// - c: The Gin context for the HTTP request.
func (h *Handler) hGetConfig(c *gin.Context) {
	ticker := strings.ToUpper(c.Param("ticker"))

	config, err := h.controller.GetConfig(ticker)
	if err != nil {
		services.WriteJSON(c, http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	shiftedConfig := config

	shiftedConfig.Drift = shiftedConfig.Drift.Mul(DRIFT_MULTIPLE)
	shiftedConfig.Volatility = shiftedConfig.Volatility.Mul(VOLATILITY_MULTIPLE)

	services.WriteJSON(c, http.StatusOK, gin.H{"ticker": ticker, "config": shiftedConfig})
}

// hPostConfig handles the HTTP POST request to update the configuration of a specific stock ticker.
//
// This endpoint accepts a JSON payload with the stock's drift and volatility values, scales them down,
// and updates the configuration in the controller.
//
// Example request payload:
//
//	{
//	    "ticker": "AAPL",
//	    "drift": 500,
//	    "volatility": 200
//	}
//
// Example response:
//
//	{
//	    "message": "success"
//	}
//
// Parameters:
// - c: The Gin context for the HTTP request.
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

	config, err := h.controller.GetConfig(payload.Ticker)
	if err != nil {
		services.WriteJSON(c, http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	config.Drift = payload.Drift.Div(DRIFT_MULTIPLE)
	config.Volatility = payload.Volatility.Div(VOLATILITY_MULTIPLE)
	h.controller.SetConfig(payload.Ticker, config)

	services.WriteJSON(c, http.StatusCreated, gin.H{"message": "success"})
}
