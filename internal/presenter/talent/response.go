package talent

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Error string `json:"error"`
}

type dataResponse struct {
	Data any `json:"data"`
}

// RespondError renders a standard error payload.
func RespondError(c *gin.Context, status int, message string) {
	c.JSON(status, errorResponse{Error: message})
}

// RespondData renders a standard data payload.
func RespondData(c *gin.Context, status int, payload any) {
	c.JSON(status, dataResponse{Data: payload})
}
