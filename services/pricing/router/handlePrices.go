package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pricing-app/services"
	"pricing-app/services/types"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

// hGetPrices handles the HTTP GET request to retrieve historical price data for a specific stock ticker.
//
// This endpoint responds with the stock's historical prices and their corresponding timestamps.
//
// Example response:
//
//	{
//	    "ticker": "AAPL",
//	    "prices": [150.25, 151.30, 152.10],
//	    "times": ["2025-03-24T09:30:00Z", "2025-03-24T09:31:00Z", "2025-03-24T09:32:00Z"]
//	}
//
// Parameters:
// - c: The Gin context for the HTTP request.
func (h *Handler) hGetPrices(c *gin.Context) {
	ticker := strings.ToUpper(c.Param("ticker"))

	exist := h.controller.CheckTickerExist(ticker)
	if !exist {
		services.WriteJSON(c, http.StatusBadRequest, types.ErrInvalidTicker)
		return
	}

	// Update the price to the latest time (minute)
	h.controller.UpdatePriceToLatestMin(ticker, time.Now())

	// Prepare response data
	prices, _ := h.controller.GetHistoryData(ticker)

	startTime, closeTime := h.controller.GetMarketTime()
	times := make([]time.Time, len(prices))
	for i, t := 0, startTime; i < len(prices) && (t.Before(closeTime) || t.Equal(closeTime)); i, t = i+1, t.Add(1*time.Minute) {
		times[i] = t
	}

	services.WriteJSON(c, http.StatusOK, gin.H{"ticker": ticker, "prices": prices, "times": times})
}

// hStreamUpdatePrice handles the HTTP GET request to stream live price updates for a specific stock ticker.
//
// This endpoint uses Server-Sent Events (SSE) to send real-time price updates to the client.
//
// Example response (SSE format):
//
//	data: {"ticker":"AAPL","time":"2025-03-24T09:30:00Z","price":150.25}
//
// Parameters:
// - c: The Gin context for the HTTP request.
func (h *Handler) hStreamUpdatePrice(c *gin.Context) {
	ticker := strings.ToUpper(c.Param("ticker"))
	if exist := h.controller.CheckTickerExist(ticker); !exist {
		msg := fmt.Sprintf("data: ticker not exist\n\n")
		_, err := fmt.Fprintf(c.Writer, msg)
		if err != nil {
			return
		}

		c.Writer.Flush()
		return
	}

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	var updateData = func(ticker string, now time.Time) decimal.Decimal {
		latestPrice, _ := h.controller.UpdatePrice(ticker, now)
		return latestPrice
	}

	var sendData = func(ticker string, now time.Time, latestPrice decimal.Decimal) {
		jsonData, err := json.Marshal(gin.H{"ticker": ticker, "time": now, "price": latestPrice})
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := fmt.Sprintf("data: %s\n\n", jsonData)
		_, err = fmt.Fprintf(c.Writer, msg)
		if err != nil {
			// client disconnect
			fmt.Printf("client disconnected: %s\n", err)
			return
		}
		c.Writer.Flush()
	}

	// Add additional 5 ms is to compensate the floating number error
	nextMinute := time.Now().Truncate(time.Minute).Add(1*time.Minute + 5*time.Millisecond)
	delay := time.Until(nextMinute)

	// Align the timer to next full minute, for example 13:21:00
	select {
	case <-time.After(delay):
		now := time.Now()
		latestPrice := updateData(ticker, now)
		sendData(ticker, now, latestPrice)

	case <-c.Request.Context().Done():
		fmt.Println("client close connection")
		return
	}

	timeTicker := time.NewTicker(1 * time.Minute)
	defer timeTicker.Stop()

	for {
		select {
		case t := <-timeTicker.C:
			latestPrice := updateData(ticker, t)
			sendData(ticker, t, latestPrice)

		case <-c.Request.Context().Done():
			fmt.Println("client close connection")
			return
		}
	}
}
