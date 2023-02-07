package articles

import (
	"fmt"

	"github.com/Floriansylvain/GohCMS/internal/api"
	"github.com/Floriansylvain/GohCMS/internal/database"
	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context) {
	id := c.Params.ByName("id")

	var articleUpdate database.DocumentUpdate
	articleUpdate.Filter = map[string]any{"titleID": id}
	c.BindJSON(&articleUpdate.Update)

	editCount, err := database.EditDocument(articlesLocation, articleUpdate)
	if err != nil {
		api.SendBadRequest(c, fmt.Sprintf(`Could not edit document(s) from DB: %v`, err.Error()))
		return
	}

	if editCount != 0 {
		api.SendOk(c, fmt.Sprintf("%d article(s) was/were successfully edited!", editCount))
	} else {
		api.SendOk(c, "No articles were edited.")
	}
}
