package pages

import (
	"GoCMS/api"
	"html/template"
	"net/http"
)

func GetPostsPage(w http.ResponseWriter, _ *http.Request) {
	posts := api.Container.ListPostsUseCase.ListPosts()

	var formattedPosts []map[string]any
	for _, post := range posts {
		formattedPosts = append(formattedPosts, map[string]any{
			"ID":        post.ID,
			"Title":     post.Title,
			"IsOnline":  post.IsOnline,
			"CreatedAt": post.CreatedAt.Format("2006-01-02 15:04:05"),
			"UpdatedAt": post.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	navbarTmpl, _ := api.Container.GetPageUseCase.GetPage("componentNavbar", nil)
	postsTmpl, _ := api.Container.GetPageUseCase.GetPage("posts", map[string]any{
		"Navbar": template.HTML(navbarTmpl),
		"Head":   headTmpl,
		"Posts":  formattedPosts,
	})

	_, _ = w.Write(postsTmpl)
}
