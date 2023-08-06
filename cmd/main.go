package main

import (
	"GohCMS2/api"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"net/http"
)

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	api.InitContainer()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		msg, _ := json.Marshal(map[string]string{"message": "Hello World"})
		_, _ = w.Write(msg)
	})
	r.Mount("/article", api.NewArticleRouter())

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
