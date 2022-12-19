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

func initBasicRoutes(r *gin.Engine) {
	r.POST("/login", internal.AuthMiddleware.LoginHandler)
	r.GET("/ping", internal.Ping)
}

func initArticlesRoutes(r *gin.Engine) {
	articlesRouter := r.Group("/articles/")
	articlesRouter.Use(internal.AuthMiddleware.MiddlewareFunc())

	articlesRouter.GET("/", internal.GetAllArticles)
	articlesRouter.GET("/:id", internal.GetArticle)
	articlesRouter.POST("/:id", internal.AddArticle)
	articlesRouter.DELETE("/:id", internal.DeleteArticle)
}

func initGin() {
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		gin.SetMode(ginMode)
	}

	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})

	initBasicRoutes(r)
	initArticlesRoutes(r)

	r.Run(":" + os.Getenv("API_PORT"))
}

func main() {
	initEnvVariables()
	initJWT()
	initGin()
}
