package pages

import (
	"GoCMS/api"
	"html/template"
	"net/http"
	"os"
)

func GetPageIntegration(w http.ResponseWriter, r *http.Request) {
	navbarTmpl, _ := api.Container.GetPageUseCase.GetPage("componentNavbar", nil)
	templ, _ := api.Container.GetPageUseCase.GetPage("integration", map[string]any{
		"Navbar": template.HTML(navbarTmpl),
		"Head":   headTmpl,
		"Host":   os.Getenv("HOST"),
	})
	_, _ = w.Write(templ)
}
