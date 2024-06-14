package api

import (
	"html/template"
	"net/http"
	"os"
)

func GetPageIntegration(w http.ResponseWriter, r *http.Request) {
	navbarTmpl, _ := Container.GetPageUseCase.GetPage("componentNavbar", nil)
	templ, _ := Container.GetPageUseCase.GetPage("integration", map[string]interface{}{
		"Navbar": template.HTML(navbarTmpl),
		"Head":   headTmpl,
		"Host":   os.Getenv("HOST"),
	})
	_, _ = w.Write(templ)
}
