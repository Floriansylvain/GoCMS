package pages

import (
	"GoCMS/api"
	"GoCMS/useCases"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

type PostCreatePageError struct {
	IsError bool
	Message string
}

func GetPostCreatePageTemplate(postName string, errorMessage string) ([]byte, error) {
	navbarTmpl, _ := api.Container.GetPageUseCase.GetPage("componentNavbar", nil)
	return api.Container.GetPageUseCase.GetPage("postCreate", map[string]interface{}{
		"Navbar": template.HTML(navbarTmpl),
		"Head":   headTmpl,
		"PageError": PostCreatePageError{
			IsError: errorMessage != "",
			Message: errorMessage,
		},
		"Name": postName,
	})
}

func PostPostCreatePage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	postName := r.FormValue("name")
	pattern := regexp.MustCompile("^[a-zA-Z0-9À-ÖØ-öø-ÿĀ-ſḀ-ỿ ]{3,50}$")

	if !pattern.MatchString(postName) {
		postsTmpl, _ := GetPostCreatePageTemplate(postName, "Name should be alphanumeric, and between 3 and 50 characters.")
		_, _ = w.Write(postsTmpl)
		return
	}

	post, err := api.Container.CreatePostUseCase.CreatePost(useCases.CreatePostCommand{
		Title: postName,
		Body:  "",
	})
	if err != nil {
		postsTmpl, _ := GetPostCreatePageTemplate(postName, "Something went wrong when creating the post, please contact admin.")
		_, _ = w.Write(postsTmpl)
		return
	}

	http.Redirect(w, r, "/post/"+strconv.Itoa(int(post.ID))+"/edit", http.StatusSeeOther)
}

func GetPostCreatePage(w http.ResponseWriter, _ *http.Request) {
	postsTmpl, _ := GetPostCreatePageTemplate("", "")
	_, _ = w.Write(postsTmpl)
}
