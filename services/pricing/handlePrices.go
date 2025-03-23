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
	currTime := getCurrentTime()
	isOpen := checkMarketOpen(currTime)

	if !isOpen {
		services.WriteJSON(c, http.StatusForbidden, gin.H{"message": "Market is closed. Available between 09:30 and 16:00."})
		return
	}

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
	closeTime := time.Date(now.Year(), now.Month(), now.Day(), MARKET_CLOSE_TIME.Hour(), MARKET_CLOSE_TIME.Minute(), MARKET_CLOSE_TIME.Second(), 0, time.Local)
	times := make([]time.Time, tail)
	for i, t := 0, startTime; i < tail && (t.Before(closeTime) || t.Equal(closeTime)); i, t = i+1, t.Add(1*time.Minute) {
		times[i] = t
	}

	services.WriteJSON(c, http.StatusOK, gin.H{"ticker": ticker, "prices": prices, "times": times})
}

func (h *Handler) hStreamUpdatePrice(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
}
