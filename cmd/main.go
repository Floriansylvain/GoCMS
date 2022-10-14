package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gotest/internal"
	"github.com/joho/godotenv"
)

func initEnvVariables() {
	if godotenv.Load() != nil {
		panic("Error loading .env file.")
	}
}

func initGin() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})

	r.GET("/ping", internal.Ping)
	r.POST("/add-article", internal.AddArticle)

	r.Run(":" + os.Getenv("API_PORT"))
}

func main() {
	initEnvVariables()
	initGin()
}
