package server

import (
	"GohCMS2/api"
	"GohCMS2/main/route"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

var possibleEnvFileLocations = []string{".env", "../.env"}
var envVarsToLoad = []string{"PORT", "ENVIRONMENT", "CORS_ALLOWED_ORIGINS"}

func initEnvVariables() {
	var err error
	for _, envLocation := range possibleEnvFileLocations {
		err = godotenv.Load(envLocation)
		if err == nil {
			break
		}
	}
	if err != nil {
		panic("Could not load .env file")
	}

	for _, envVar := range envVarsToLoad {
		if _, ok := os.LookupEnv(envVar); !ok {
			panic(fmt.Sprintf("Environment variable %s is not set", envVar))
		}
	}
}

func InitServer() *chi.Mux {
	initEnvVariables()
	api.InitContainer()
	api.InitValidator()
	route.InitJwt()
	return route.InitRoutes()
}

func StartServer(router *chi.Mux) error {
	fmt.Println("Server starting on http://localhost:" + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
	return err
}
