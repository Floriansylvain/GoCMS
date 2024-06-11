package api

import (
	"html/template"
	"net/http"
)

func GetPostsPage(w http.ResponseWriter, _ *http.Request) {
	posts := Container.ListPostsUseCase.ListPosts()

	var formattedPosts []map[string]interface{}
	for _, post := range posts {
		formattedPosts = append(formattedPosts, map[string]interface{}{
			"ID":        post.ID,
			"Title":     post.Title,
			"CreatedAt": post.CreatedAt.Format("2006-01-02 15:04:05"),
			"UpdatedAt": post.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	navbarTmpl, _ := Container.GetPageUseCase.GetPage("componentNavbar", nil)
	postsTmpl, _ := Container.GetPageUseCase.GetPage("posts", map[string]interface{}{
		"Navbar": template.HTML(navbarTmpl),
		"Head":   headTmpl,
		"Posts":  formattedPosts,
	})

	_, _ = w.Write(postsTmpl)
}
