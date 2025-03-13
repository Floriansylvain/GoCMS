package post

import (
	"GoCMS/api"
	"GoCMS/api/controllers/auth"
	"GoCMS/domain/post"
	"GoCMS/useCases"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PostPost struct {
	Title string `json:"title" validate:"required,min=3,max=50"`
	Body  string `json:"body" validate:"required,max=10000"`
}

const IdUint32ErrorMessage = "The server expects the ID to be in the format of an unsigned 32-bit integer (uint32)."

func getPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, IdUint32ErrorMessage, http.StatusBadRequest)
		return
	}

	localPost, err := api.Container.GetPostUseCase.GetPost(uint32(id))
	if err != nil || !localPost.IsOnline {
		http.Error(w, "The requested resource, identified by its unique ID, could not be found on the server.", http.StatusNotFound)
		return
	}

	postJson, _ := json.Marshal(localPost)
	_, _ = w.Write(postJson)
}

func postPost(w http.ResponseWriter, r *http.Request) {
	var localPost PostPost
	err := json.NewDecoder(r.Body).Decode(&localPost)
	if err != nil {
		http.Error(w, auth.BodyErrorMessage, http.StatusBadRequest)
		return
	}

	err = api.Validate.Struct(localPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdPost, err := api.Container.CreatePostUseCase.CreatePost(useCases.CreatePostCommand{
		Title: localPost.Title,
		Body:  localPost.Body,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	postJson, _ := json.Marshal(createdPost)

	_, _ = w.Write(postJson)
}

func listPosts(w http.ResponseWriter, _ *http.Request) {
	posts := api.Container.ListPostsUseCase.ListPosts()
	onlinePosts := make([]post.Post, 0)
	for _, localPost := range posts {
		if localPost.IsOnline {
			onlinePosts = append(onlinePosts, localPost)
		}
	}
	postsJson, _ := json.Marshal(onlinePosts)
	_, _ = w.Write(postsJson)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, IdUint32ErrorMessage, http.StatusBadRequest)
	}

	err = api.Container.DeletePostUseCase.DeletePost(uint32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	_, _ = w.Write([]byte("post deleted"))
}

func NewPostRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", getPost)
	r.Post("/", postPost)
	r.Get("/", listPosts)
	r.Delete("/{id}", deletePost)
	return r
}
