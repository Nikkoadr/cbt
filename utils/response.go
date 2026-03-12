package utils

import "github.com/gin-gonic/gin"

// APIResponse represents a standard API response
type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SendSuccess sends a success response
func SendSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(200, APIResponse{Message: message, Data: data})
}

// SendError sends an error response
func SendError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, APIResponse{Message: message})
}
