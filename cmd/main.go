package main

import (
	"GohCMS2/api"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
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

func main() {
	api.InitContainer()
	api.InitValidator()
	initJwt()

	r := chi.NewRouter()
	r.Get("/", getHelloWorld)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(api.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Mount("/article", api.NewArticleRouter())
	})

	r.Mount("/auth", api.NewAuthRouter())

	apiRouter := chi.NewRouter()
	apiRouter.Use(httplog.RequestLogger(httplog.NewLogger("GohCMS2")))
	apiRouter.Use(jsonContentTypeMiddleware)
	apiRouter.Mount("/v1", r)

	fmt.Println("Server starting on port 8080")
	err := http.ListenAndServe(":8080", apiRouter)
	if err != nil {
		panic(err)
	}
}
