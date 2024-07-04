package commons

import (
	"github.com/gin-gonic/gin"
)

type baseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func Response(c *gin.Context, data interface{}, message string, status int, err string) {
	c.JSON(status, baseResponse{
		Message: message,
		Data:    data,
		Error:   err,
	})
}
