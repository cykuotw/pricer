package router

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// hStreamServerTime handles the HTTP GET request to stream the current server time to the client.
//
// This endpoint uses Server-Sent Events (SSE) to send the current server time every second.
//
// Example response (SSE format):
//
//	data: 2025-03-24T09:30:00.123456789Z
//
// Parameters:
// - c: The Gin context for the HTTP request.
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
