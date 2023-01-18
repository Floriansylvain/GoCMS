package articles

import (
	"fmt"

	"github.com/Floriansylvain/GohCMS/internal/api"
	"github.com/Floriansylvain/GohCMS/internal/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Add(c *gin.Context) {
	var article Article
	article.TitleID = c.Params.ByName("id")
	if c.BindJSON(&article) != nil {
		api.SendBadRequest(c, "Could not correctly parse the article.")
		return
	}

	document, err := bson.Marshal(article)
	if err != nil {
		api.SendBadRequest(c, "Could not correctly marshal the article.")
		return
	}

	documents, _ := database.GetDocuments(articlesLocation, map[string]any{})
	if IsArticleIdAlreadyUsed(article.TitleID, documents) {
		api.SendBadRequest(c, fmt.Sprintf("Article ID '%v' already used.", article.TitleID))
		return
	}

	err = database.PushDocument(articlesLocation, document)
	if err != nil {
		api.SendBadRequest(c, fmt.Sprintf(`Could not insert document(s) into DB: %v`, err.Error()))
		return
	}

	api.SendOk(c, "Article successfully added!")
}
