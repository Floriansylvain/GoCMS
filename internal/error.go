package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorMessage struct {
	Message string `json:"message"`
}

func SendErrorMessageToClient(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, errorMessage{Message: message})
}
