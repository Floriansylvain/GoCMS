package articles

import (
	"fmt"

	"github.com/Floriansylvain/GohCMS/internal/api"
	"github.com/Floriansylvain/GohCMS/internal/database"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	deleteCount, err := database.DeleteDocument(articlesLocation, map[string]any{"titleID": id})
	if err != nil {
		api.SendBadRequest(c, fmt.Sprintf(`Could not delete document(s) from DB: %v`, err.Error()))
		return
	}

	if deleteCount != 0 {
		api.SendOk(c, fmt.Sprintf("%d article(s) was/were successfully deleted!", deleteCount))
	} else {
		api.SendOk(c, "No articles were deleted.")
	}
}
