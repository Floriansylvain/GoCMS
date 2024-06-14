package api

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func updateIsOnline(postId string, isOnline bool) (error, int) {
	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		return errors.New("the server expects the ID to be in the format of an unsigned 32-bit integer (uint32)"), http.StatusBadRequest
	}

	_, err = Container.UpdatePostUseCase.UpdateIsOnline(uint32(postIdInt), isOnline)
	if err != nil {
		return errors.New("the requested resource, identified by its unique ID, could not be found on the server"), http.StatusNotFound
	}

	return nil, http.StatusOK
}

func GetPostUnpublishPage(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")
	err, statusCode := updateIsOnline(postId, false)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	http.Redirect(w, r, "/post", http.StatusSeeOther)
}

func GetPostPublishPage(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")
	err, statusCode := updateIsOnline(postId, true)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	http.Redirect(w, r, "/post", http.StatusSeeOther)
}
