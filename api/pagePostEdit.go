package api

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"strconv"
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
	postID := chi.URLParam(r, "id")
	postIDint, err := strconv.Atoi(postID)
	if err != nil {
		_, _ = w.Write(getPostEditPageTemplate("", PostEditPageAlert{
			IsError: true,
			Message: "Could not find the requested post.",
		}))
		return
	}

	_ = r.ParseForm()
	postBody := r.FormValue("postBody")

	post, _ := Container.GetPostUseCase.GetPost(uint32(postIDint))
	err = Container.UpdatePostUseCase.UpdateBody(post.ID, postBody)
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
	postID := chi.URLParam(r, "id")
	postIDint, err := strconv.Atoi(postID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post, _ := Container.GetPostUseCase.GetPost(uint32(postIDint))

	_, _ = w.Write(getPostEditPageTemplate(post.Body, PostEditPageAlert{
		IsError: false,
		Message: "",
	}))
}
