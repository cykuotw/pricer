package services

import (
	"encoding/json"
	"io"
	"pricing-app/services/types"

	"github.com/gin-gonic/gin"
)

/*
	Private utility functions within package services
*/

// WriteJSON writes a JSON response to the client.
//
// Parameters:
// - c: The Gin context for the HTTP request.
// - status: The HTTP status code to send in the response.
// - obj: The object to encode as JSON and send in the response.
func WriteJSON(c *gin.Context, status int, obj any) {
	c.Header("Content-Type", "application/json")
	c.JSON(status, obj)
}

// ParseJSON parses the JSON body of an HTTP request into the provided payload.
//
// Parameters:
// - c: The Gin context for the HTTP request.
// - payload: A pointer to the object where the parsed JSON data will be stored.
//
// Returns:
// - An error if the request body is empty or if the JSON decoding fails.
func ParseJSON(c *gin.Context, payload any) error {
	err := json.NewDecoder(c.Request.Body).Decode(payload)
	if err == io.EOF {
		return types.ErrEmptyRequestBody
	}
	if err != nil {
		return err
	}
	return nil
}
