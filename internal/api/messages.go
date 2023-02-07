package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"message": message, "code": 400})
}

func SendOk(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{"message": message, "code": 200})
}

func SendForbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, gin.H{"message": message, "code": 403})
}
