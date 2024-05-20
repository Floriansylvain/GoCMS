package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func GetPostDeletePage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, idUint32ErrorMessage, http.StatusBadRequest)
		return
	}

	err = Container.DeletePostUseCase.DeletePost(uint32(id))
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/post", http.StatusSeeOther)
}
