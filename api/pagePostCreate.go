package api

import (
	"html/template"
	"net/http"
	"regexp"
)

func PostPostCreatePage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	postName := r.FormValue("name")
	pattern := regexp.MustCompile("^[a-zA-Z0-9]{3,50}$")

	isPostNameValid := pattern.MatchString(postName)
	if !isPostNameValid {
		r.Method = http.MethodGet

	}
}

func GetPostCreatePage(w http.ResponseWriter, _ *http.Request) {
	navbarTmpl, _ := Container.GetPageUseCase.GetPage("componentNavbar", nil)
	postsTmpl, _ := Container.GetPageUseCase.GetPage("postCreate", map[string]interface{}{
		"Navbar": template.HTML(navbarTmpl),
		"Head":   headTmpl,
	})
	_, _ = w.Write(postsTmpl)
}
