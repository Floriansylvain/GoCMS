package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func GetPostDeletePage(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(postId)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	err = Container.DeletePostUseCase.DeletePost(uint32(id))
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/post", http.StatusSeeOther)
}
