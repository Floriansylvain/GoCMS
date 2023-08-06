package api

import (
	. "GohCMS2/domain/article"
	. "GohCMS2/useCases"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func getArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, "The server expects the ID to be in the format of an unsigned 32-bit integer (uint32).", http.StatusBadRequest)
		return
	}

	article, err := container.GetArticleUseCase.GetArticle(uint32(id))
	if err != nil {
		http.Error(w, "The requested resource, identified by its unique ID, could not be found on the server.", http.StatusNotFound)
		return
	}

	articleJson, _ := json.Marshal(article)
	_, _ = w.Write(articleJson)
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "The request cannot be processed due to a mismatch in the format of the body.", http.StatusBadRequest)
		return
	}

	createdArticle, err := container.CreateArticleUseCase.CreateArticle(CreateArticleCommand{
		Title: article.Title,
		Body:  article.Body,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	articleJson, _ := json.Marshal(createdArticle)

	_, _ = w.Write(articleJson)
}

func listArticles(w http.ResponseWriter, _ *http.Request) {
	articles := container.ListArticlesUseCase.ListArticles()
	articlesJson, _ := json.Marshal(articles)

	_, _ = w.Write(articlesJson)
}

func NewArticleRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", getArticle)
	r.Post("/", postArticle)
	r.Get("/", listArticles)
	return r
}
