package api

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
)

func GetPostEditPage(w http.ResponseWriter, r *http.Request) {
	postName := chi.URLParam(r, "name")
	if len(postName) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post, _ := Container.GetPostUseCase.GetPostByName(postName)

	navbarTmpl, _ := Container.GetPageUseCase.GetPage("componentNavbar", nil)
	postsTmpl, _ := Container.GetPageUseCase.GetPage("postEdit", map[string]interface{}{
		"Navbar": template.HTML(navbarTmpl),
		"Head":   headTmpl,
		"Body":   post.Body,
	})
	_, _ = w.Write(postsTmpl)
}
