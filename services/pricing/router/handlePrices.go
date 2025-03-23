package router

import (
	"fmt"
	"net/http"
	"pricing-app/services"
	"pricing-app/services/types"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) hGetPrices(c *gin.Context) {
	now := time.Now()
	isOpen := h.controller.CheckMarketOpen(now)

	if !isOpen {
		services.WriteJSON(c, http.StatusForbidden, gin.H{"message": "Market is closed. Available between 09:30 and 16:00."})
		return
	}

	ticker := strings.ToUpper(c.Param("ticker"))

	exist := h.controller.CheckTickerExist(ticker)
	if !exist {
		services.WriteJSON(c, http.StatusBadRequest, types.ErrInvalidTicker)
		return
	}

	// update the price to the latest time (minute)
	h.controller.UpdatePriceToLatestMin(ticker, now)

	// prepare response data
	prices, _ := h.controller.GetHistoryData(ticker)

	startTime, closeTime := h.controller.GetMarketTime()
	times := make([]time.Time, len(prices))
	for i, t := 0, startTime; i < len(prices) && (t.Before(closeTime) || t.Equal(closeTime)); i, t = i+1, t.Add(1*time.Minute) {
		times[i] = t
	}

	services.WriteJSON(c, http.StatusOK, gin.H{"ticker": ticker, "prices": prices, "times": times})
}

func (h *Handler) hStreamUpdatePrice(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	var sendData = func(t time.Time) {
		fmt.Printf("%s\n", t.Format(time.RFC3339Nano))
		msg := fmt.Sprintf("data: %s\n\n", time.Now().Format(time.RFC3339Nano))
		_, err := fmt.Fprintf(c.Writer, msg)
		if err != nil {
			// client disconnect
			fmt.Printf("client disconnected: %s\n", err)
			return
		}
		c.Writer.Flush()
	}

	// add additional 5 ms is to compensate the floating number error
	nextMinute := time.Now().Truncate(time.Minute).Add(1*time.Minute + 5*time.Millisecond)
	delay := time.Until(nextMinute)

	// align the timer to next full minute, for example 13:21:00
	select {
	case <-time.After(delay):
		sendData(time.Now())

	case <-c.Request.Context().Done():
		fmt.Println("client close connection")
		return
	}

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			sendData(t)

		case <-c.Request.Context().Done():
			fmt.Println("client close connection")
			return
		}
	}
}
