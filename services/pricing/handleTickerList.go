package pricing

import (
	"net/http"
	"pricing-app/services"
	"sort"

	"github.com/gin-gonic/gin"
)

func (h *Handler) hGetTickerList(c *gin.Context) {
	tickers := make([]string, 0, len(h.controller.config))
	for ticker := range h.controller.config {
		tickers = append(tickers, ticker)
	}
	sort.Strings(tickers)

	services.WriteJSON(c, http.StatusOK, gin.H{"tickers": tickers})
}
