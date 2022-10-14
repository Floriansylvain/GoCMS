package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type answer struct {
	Message string `json:"message"`
}

func SendErrorMessageToClient(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, answer{Message: message})
}

func SendOkMessageToClient(c *gin.Context, message string) {
	c.JSON(http.StatusOK, answer{Message: message})
}
