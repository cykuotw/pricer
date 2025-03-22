package pricing

import (
	"net/http"
	"pricing-app/services"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func (h *Handler) hGetPrices(c *gin.Context) {
	ticker := strings.ToUpper(c.Param("ticker"))

	_, ok := h.controller.config[ticker]
	if !ok {
		services.WriteJSON(c, http.StatusBadRequest, gin.H{"message": "invalid ticker"})
		return
	}

	tail := h.controller.historyBufferTail[ticker]
	prices := make([]decimal.Decimal, tail)
	copy(prices, h.controller.historyBuffer[ticker][:tail])

	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), MARKET_OPEN_TIME.Hour(), MARKET_OPEN_TIME.Minute(), MARKET_OPEN_TIME.Second(), 0, time.Local)

	services.WriteJSON(c, http.StatusOK, gin.H{"ticker": ticker, "prices": prices, "t0": startTime})
}

func (h *Handler) hStreamUpdatePrice(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
}
