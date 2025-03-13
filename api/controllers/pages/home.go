package pages

import (
	"GoCMS/api"
	"html/template"
	"net/http"
)

func GetHomePage(w http.ResponseWriter, _ *http.Request) {
	navbarTmpl, _ := api.Container.GetPageUseCase.GetPage("componentNavbar", nil)
	homeTmpl, _ := api.Container.GetPageUseCase.GetPage("home", map[string]interface{}{
		"Navbar": template.HTML(navbarTmpl),
		"Head":   headTmpl,
	})
	_, _ = w.Write(homeTmpl)
}
