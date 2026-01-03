package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping responds with a simple JSON payload for health checks.
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
