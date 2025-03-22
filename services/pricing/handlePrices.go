package pricing

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) hStreamUpdatePrice(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
}
