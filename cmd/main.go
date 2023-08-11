package main

import (
	"GohCMS2/api"
	"encoding/json"
	"fmt"
	"github.com/MadAppGang/httplog"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

var envVarsToLoad = []string{"PORT", "ENVIRONMENT", "CORS_ALLOWED_ORIGINS"}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func htmlContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		next.ServeHTTP(w, r)
	})
}

func initJwt() {
	api.TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func getHelloWorld(w http.ResponseWriter, _ *http.Request) {
	msg, _ := json.Marshal(map[string]string{"message": "Hello World"})
	_, _ = w.Write(msg)
}

func initBackendRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(httplog.LoggerWithName("backend"))
	r.Use(jsonContentTypeMiddleware)
	r.Get("/", getHelloWorld)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(api.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Mount("/article", api.NewArticleRouter())
	})
	r.Mount("/auth", api.NewAuthRouter())

	return r
}

func initFrontendRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(httplog.LoggerWithName("frontend"))
	r.Use(htmlContentTypeMiddleware)
	r.Mount("/", api.NewPageRouter())

	return r
}

func initRoutes() *chi.Mux {
	backend := initBackendRoutes()
	frontend := initFrontendRoutes()

	apiRouter := chi.NewRouter()
	apiRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";"),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	apiRouter.Mount("/v1", backend)
	apiRouter.Mount("/", frontend)

	return apiRouter
}

func initEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	for _, envVar := range envVarsToLoad {
		if _, ok := os.LookupEnv(envVar); !ok {
			panic(fmt.Sprintf("Environment variable %s is not set", envVar))
		}
	}
}

func main() {
	initEnvVariables()
	api.InitContainer()
	api.InitValidator()
	initJwt()
	router := initRoutes()

	fmt.Println("Server starting on http://localhost:" + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
	if err != nil {
		panic(err)
	}
}
