package articles

import (
	"fmt"
	"net/http"

	"github.com/Floriansylvain/GohCMS/internal/api"
	"github.com/Floriansylvain/GohCMS/internal/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllArticlesHandler(c *gin.Context) {
	documents, err := database.GetDocuments(articlesLocation, gin.H{})
	if err != nil {
		api.SendBadRequest(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetAllArticlesBusiness(documents))
}

func GetArticleHandler(c *gin.Context) {
	articleID := c.Params.ByName("id")
	article, err := database.GetUniqueDocument(articlesLocation,
		gin.H{"id_name": articleID})
	if err != nil {
		api.SendBadRequest(c, "The ID provided doesn't match any article.")
		return
	}
	var parsedArticle Article
	bson.Unmarshal(article, &parsedArticle)
	c.JSON(http.StatusOK, parsedArticle)
}

func AddArticleHandler(c *gin.Context) {
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

	documents, _ := database.GetDocuments(articlesLocation, gin.H{})
	if IsArticleIdAlreadyUsed(article.TitleID, documents) {
		api.SendBadRequest(c, "Article ID already used.")
		return
	}

	err = database.PushDocument(articlesLocation, document)
	if err != nil {
		api.SendBadRequest(c, fmt.Sprintf(`Could not insert document(s) into DB: %v`, err.Error()))
		return
	}

	api.SendOk(c, "Article successfully added!")
}

func DeleteArticleHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	deleteCount, err := database.DeleteDocument(articlesLocation, gin.H{"id_name": id})
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

func EditArticleHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	var articleUpdate database.DocumentUpdate
	articleUpdate.Filter = gin.H{"id_name": id}
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
