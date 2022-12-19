package main

import (
	"fmt"
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

func initJWT() {
	errInit := internal.AuthMiddleware.MiddlewareInit()
	if errInit != nil {
		fmt.Printf(errInit.Error())
	}
}

func initGin() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})

	r.POST("/login", internal.AuthMiddleware.LoginHandler)
	r.GET("/ping", internal.Ping)

	articlesRouter := r.Group("/articles/")
	articlesRouter.Use(internal.AuthMiddleware.MiddlewareFunc())

	articlesRouter.GET("/", internal.GetAllArticles)
	articlesRouter.GET("/:id", internal.GetArticle)
	articlesRouter.POST("/:id", internal.AddArticle)
	articlesRouter.DELETE("/:id", internal.DeleteArticle)

	r.Run(":" + os.Getenv("API_PORT"))
}

func main() {
	initEnvVariables()
	initJWT()
	initGin()
}
