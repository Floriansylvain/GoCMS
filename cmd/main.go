package main

import (
	"fmt"
	"os"

	"github.com/Floriansylvain/GohCMS/internal/api"
	"github.com/Floriansylvain/GohCMS/internal/articles"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var ginMode string
var apiPort string
var frontPort string
var hostAddress string

func initEnvVariables() {
	if godotenv.Load() != nil {
		panic("Error loading .env file.")
	}

	ginMode = os.Getenv("APP_GIN_MODE")
	apiPort = os.Getenv("APP_API_PORT")
	frontPort = os.Getenv("APP_FRONT_PORT")
	hostAddress = os.Getenv("APP_HOST_ADDRESS")
}

func initJWT() {
	errInit := api.AuthMiddleware.MiddlewareInit()
	if errInit != nil {
		fmt.Printf(errInit.Error())
	}
}

func initBasicRoutes(r *gin.Engine) {
	r.POST("/login/", api.AuthMiddleware.LoginHandler)
	r.GET("/ping/", api.Ping)
}

func corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("http://%v:%v", hostAddress, frontPort))
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
	articlesRouter.Use(corsMiddleware, api.AuthMiddleware.MiddlewareFunc())

	articlesRouter.GET("/", articles.GetArticleHandler)
	articlesRouter.GET("/:id", articles.GetArticleHandler)
	articlesRouter.POST("/:id", articles.AddArticleHandler)
	articlesRouter.PATCH("/:id", articles.EditArticleHandler)
	articlesRouter.DELETE("/:id", articles.DeleteArticleHandler)
}

func initGin() {
	r := gin.Default()
	r.Use(corsMiddleware)

	if ginMode == "release" {
		gin.SetMode(ginMode)
	}

	initBasicRoutes(r)
	initArticlesRoutes(r)

	r.Run(":" + apiPort)
}

func main() {
	initEnvVariables()
	initJWT()
	initGin()
}
