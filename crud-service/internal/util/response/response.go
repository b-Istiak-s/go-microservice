package response

import (
	"github.com/gin-gonic/gin"
)

// Success sends a standard success JSON response.
func Success(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

// Error sends a standard error JSON response.
func Error(c *gin.Context, status int, message string, err ...interface{}) {
	c.JSON(status, gin.H{
		"success": false,
		"message": message,
		"error":   err,
	})
}
