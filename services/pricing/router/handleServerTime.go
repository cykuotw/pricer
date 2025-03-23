package router

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) hStreamServerTime(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			msg := fmt.Sprintf("data: %s\n\n", t.Format(time.RFC3339Nano))
			_, err := fmt.Fprintf(c.Writer, msg)
			if err != nil {
				// client disconnect
				fmt.Printf("client disconnected: %s\n", err)
				return
			}
			c.Writer.Flush()

		case <-c.Request.Context().Done():
			fmt.Println("client close connection")
			return
		}
	}
}
