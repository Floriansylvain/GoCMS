package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gotest/internal"
	"github.com/joho/godotenv"
)

var ginMode = os.Getenv("APP_GIN_MODE")

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
	r.POST("/login/", internal.AuthMiddleware.LoginHandler)
	r.GET("/ping/", internal.Ping)
}

func corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

func initArticlesRoutes(r *gin.Engine) {
	articlesRouter := r.Group("/articles")
	articlesRouter.Use(internal.AuthMiddleware.MiddlewareFunc())
	if ginMode != "release" {
		articlesRouter.Use(corsMiddleware)
	}

	articlesRouter.GET("/", internal.GetAllArticlesHandler)
	articlesRouter.GET("/:id", internal.GetArticleHandler)
	articlesRouter.POST("/:id", internal.AddArticleHandler)
	articlesRouter.DELETE("/:id", internal.DeleteArticleHandler)
}

func initGin() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})

	if ginMode == "release" {
		gin.SetMode(ginMode)
	} else {
		r.Use(corsMiddleware)
	}

	initBasicRoutes(r)
	initArticlesRoutes(r)

	r.Run(":" + os.Getenv("APP_API_PORT"))
}

func main() {
	initEnvVariables()
	initJWT()
	initGin()
}
