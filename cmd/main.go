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

	r.GET("/get-all-articles", internal.AuthCheck, internal.GetAllArticles)
	r.GET("/articles/:id", internal.AuthCheck, internal.GetArticle)
	r.POST("/add-article", internal.AuthCheck, internal.AddArticle)
	r.DELETE("/delete-article", internal.AuthCheck, internal.DeleteArticle)

	r.POST("/login", internal.LoginUser)
	r.POST("/logout", internal.LogoutUser)

	r.Run(":" + os.Getenv("API_PORT"))
}

func main() {
	initEnvVariables()
	initGin()
}
