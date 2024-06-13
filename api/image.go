package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func postImage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newImage, err := Container.CreateImageUseCase.CreateImage(file, *fileHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Container.UpdatePostUseCase.AddImage(uint32(idInt), newImage.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newJson := map[string]interface{}{"location": newImage.Path}
	newJsonBytes, _ := json.Marshal(newJson)

	_, _ = w.Write(newJsonBytes)
}
