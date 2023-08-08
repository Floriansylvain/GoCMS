package main

import (
	"GohCMS2/api"
	"encoding/json"
	"fmt"
	"github.com/MadAppGang/httplog"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

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
	// TODO use env variable for allowed origins
	apiRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
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

func main() {
	api.InitContainer()
	api.InitValidator()
	initJwt()

	router := initRoutes()

	fmt.Println("Server starting on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
