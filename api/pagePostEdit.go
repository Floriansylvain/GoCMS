package api

import (
	"html/template"
	"net/http"
)

func GetPostEditPage(w http.ResponseWriter, _ *http.Request) {
	navbarTmpl, _ := Container.GetPageUseCase.GetPage("componentNavbar", nil)
	postsTmpl, _ := Container.GetPageUseCase.GetPage("postEdit", map[string]interface{}{
		"Navbar": template.HTML(navbarTmpl),
		"Head":   headTmpl,
	})
	_, _ = w.Write(postsTmpl)
}
