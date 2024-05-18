package api

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"net/url"
)

type PostEditPageAlert struct {
	IsError bool
	Message string
}

func getPostEditPageTemplate(body string, alert PostEditPageAlert) []byte {
	navbarTmpl, _ := Container.GetPageUseCase.GetPage("componentNavbar", nil)
	postTmpl, _ := Container.GetPageUseCase.GetPage("postEdit", map[string]interface{}{
		"Navbar": template.HTML(navbarTmpl),
		"Head":   headTmpl,
		"Body":   body,
		"Alert":  alert,
	})
	return postTmpl
}

func PostPostEditPage(w http.ResponseWriter, r *http.Request) {
	postName := chi.URLParam(r, "name")
	if len(postName) == 0 {
		_, _ = w.Write(getPostEditPageTemplate("", PostEditPageAlert{
			IsError: true,
			Message: "Could not find the requested post.",
		}))
		return
	}

	_ = r.ParseForm()
	postBody := r.FormValue("postBody")

	parsedName, _ := url.PathUnescape(postName)
	post, _ := Container.GetPostUseCase.GetPostByName(parsedName)
	err := Container.UpdatePostUseCase.UpdateBody(post.ID, postBody)
	if err != nil {
		_, _ = w.Write(getPostEditPageTemplate(postBody, PostEditPageAlert{
			IsError: true,
			Message: "Could not save the post: " + err.Error(),
		}))
		return
	}

	_, _ = w.Write(getPostEditPageTemplate(postBody, PostEditPageAlert{
		IsError: false,
		Message: "Post successfully edited!",
	}))
}

func GetPostEditPage(w http.ResponseWriter, r *http.Request) {
	postName := chi.URLParam(r, "name")
	if len(postName) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	parsedName, _ := url.PathUnescape(postName)
	post, _ := Container.GetPostUseCase.GetPostByName(parsedName)
	_, _ = w.Write(getPostEditPageTemplate(post.Body, PostEditPageAlert{
		IsError: false,
		Message: "",
	}))
}
