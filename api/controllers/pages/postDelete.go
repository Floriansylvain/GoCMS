package pages

import (
	"GoCMS/api"
	"GoCMS/api/controllers/post"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetPostDeletePage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, post.IdUint32ErrorMessage, http.StatusBadRequest)
		return
	}

	localPost, err := api.Container.GetPostUseCase.GetPost(uint32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Println(localPost.Images)
	for _, image := range localPost.Images {
		err = api.Container.DeleteImageUseCase.DeleteImage(image.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	err = api.Container.DeletePostUseCase.DeletePost(uint32(id))
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/post", http.StatusSeeOther)
}
