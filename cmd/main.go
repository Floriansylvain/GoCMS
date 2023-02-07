package main

import (
	"fmt"
	"os"

	"github.com/Floriansylvain/GohCMS/internal/api"
	"github.com/Floriansylvain/GohCMS/internal/articles"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	ginMode      string
	apiPort      string
	apiBaseUrl   string
	frontAddress string
)

func initEnvVariables() {
	if godotenv.Load() != nil {
		panic("Error loading .env file.")
	}

	ginMode = os.Getenv("APP_GIN_MODE")
	apiPort = os.Getenv("APP_API_PORT")
	apiBaseUrl = os.Getenv("APP_BASE_API_PATH")
	frontAddress = os.Getenv("APP_FRONT_ADDRESS")
}

func initJWT() {
	errInit := api.AuthMiddleware.MiddlewareInit()
	if errInit != nil {
		fmt.Printf(errInit.Error())
	}
}

func initBasicRoutes(r *gin.RouterGroup) {
	r.POST("/login/", api.AuthMiddleware.LoginHandler)
	r.POST("/logout/", api.AuthMiddleware.LogoutHandler)
	r.GET("/ping/", api.Ping)
}

func corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", frontAddress)
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

func jwtProxyMiddleware(c *gin.Context) {
	jwtToken, _ := c.Cookie("jwt")
	c.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
	c.Next()
}

func initArticlesRoutes(r *gin.RouterGroup) {
	articlesRouter := r.Group("/articles")
	articlesRouter.Use(corsMiddleware, api.AuthMiddleware.MiddlewareFunc())

	articlesRouter.GET("/", articles.Get)
	articlesRouter.GET("/:id", articles.GetUnique)
	articlesRouter.POST("/:id", articles.Add)
	articlesRouter.PUT("/:id", articles.Edit)
	articlesRouter.DELETE("/:id", articles.Delete)
}

func initGin() {
	r := gin.Default()
	r.Use(jwtProxyMiddleware, corsMiddleware)

	var router *gin.RouterGroup
	if apiBaseUrl != "" {
		router = r.Group(apiBaseUrl)
	} else {
		router = &r.RouterGroup
	}

	if ginMode == "release" {
		gin.SetMode(ginMode)
		api.AuthMiddleware.SecureCookie = true
	}

	initBasicRoutes(router)
	initArticlesRoutes(router)

	r.Run(":" + apiPort)
}

func main() {
	initEnvVariables()
	initGin()
	initJWT()
}
