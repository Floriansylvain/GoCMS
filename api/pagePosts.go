package api

import (
	"html/template"
	"net/http"
)

func GetPostsPage(w http.ResponseWriter, _ *http.Request) {
	navbarTmpl, _ := Container.GetPageUseCase.GetPage("componentNavbar", nil)
	postsTmpl, _ := Container.GetPageUseCase.GetPage("posts", map[string]interface{}{
		"Navbar": template.HTML(navbarTmpl),
	})
	_, _ = w.Write(postsTmpl)
}
