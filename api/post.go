package api

import (
	"GohCMS2/useCases"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type PostPost struct {
	Title string `json:"title" validate:"required,min=3,max=50"`
	Body  string `json:"body" validate:"required,max=10000"`
}

func getPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, "The server expects the ID to be in the format of an unsigned 32-bit integer (uint32).", http.StatusBadRequest)
		return
	}

	post, err := Container.GetArticleUseCase.GetArticle(uint32(id))
	if err != nil {
		http.Error(w, "The requested resource, identified by its unique ID, could not be found on the server.", http.StatusNotFound)
		return
	}

	postJson, _ := json.Marshal(post)
	_, _ = w.Write(postJson)
}

func postPost(w http.ResponseWriter, r *http.Request) {
	var post PostPost
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, bodyErrorMessage, http.StatusBadRequest)
		return
	}

	err = validate.Struct(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdArticle, err := Container.CreateArticleUseCase.CreatePost(useCases.CreatePostCommand{
		Title: post.Title,
		Body:  post.Body,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	postJson, _ := json.Marshal(createdArticle)

	_, _ = w.Write(postJson)
}

func listPosts(w http.ResponseWriter, _ *http.Request) {
	posts := Container.ListArticlesUseCase.ListArticles()
	postsJson, _ := json.Marshal(posts)

	_, _ = w.Write(postsJson)
}

func NewArticleRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", getPost)
	r.Post("/", postPost)
	r.Get("/", listPosts)
	return r
}
