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
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	article := container.GetArticleUseCase.GetArticle(id)
	articleJson, err := json.Marshal(article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, _ = w.Write(articleJson)
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdArticle := container.CreateArticleUseCase.CreateAarticle(CreateArticleCommand{
		Title: article.Title,
		Body:  article.Body,
	})
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
