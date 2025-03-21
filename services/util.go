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

func WriteJSON(c *gin.Context, status int, obj any) {
	c.Header("Content-Type", "application/json")
	c.JSON(status, obj)
}

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
