package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"message": message})
}

func SendOk(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func SendForbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, gin.H{"message": message})
}
