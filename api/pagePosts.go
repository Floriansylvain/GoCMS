package api

import (
	"html/template"
	"net/http"
)

func GetPostsPage(w http.ResponseWriter, _ *http.Request) {
	posts := Container.ListPostsUseCase.ListPosts()
	navbarTmpl, _ := Container.GetPageUseCase.GetPage("componentNavbar", nil)
	postsTmpl, _ := Container.GetPageUseCase.GetPage("posts", map[string]interface{}{
		"Navbar": template.HTML(navbarTmpl),
		"Head":   headTmpl,
		"Posts":  posts,
	})
	_, _ = w.Write(postsTmpl)
}
