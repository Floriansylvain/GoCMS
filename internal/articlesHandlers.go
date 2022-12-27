package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllArticlesHandler(c *gin.Context) {
	documents, err := getDocuments(articlesLocation, gin.H{})
	if err != nil {
		SendBadRequest(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetAllArticlesBusiness(documents))
}

func GetArticleHandler(c *gin.Context) {
	articleID := c.Params.ByName("id")
	article, err := getUniqueDocument(articlesLocation,
		gin.H{"id_name": articleID})
	if err != nil {
		SendBadRequest(c, "The ID provided doesn't match any article.")
		return
	}
	var parsedArticle Article
	bson.Unmarshal(article, &parsedArticle)
	c.JSON(http.StatusOK, parsedArticle)
}

func AddArticleHandler(c *gin.Context) {
	var article Article
	article.IdName = c.Params.ByName("id")
	if c.BindJSON(&article) != nil {
		SendBadRequest(c, "Could not correctly parse the article.")
		return
	}

	document, err := bson.Marshal(article)
	if err != nil {
		SendBadRequest(c, "Could not correctly marshal the article.")
		return
	}

	documents, _ := getDocuments(articlesLocation, gin.H{})
	if IsArticleIdAlreadyUsed(article.IdName, documents) {
		SendBadRequest(c, "Article ID already used.")
		return
	}

	err = pushDocument(articlesLocation, document)
	if err != nil {
		SendBadRequest(c, fmt.Sprintf(`Could not insert document into DB: %v`, err.Error()))
		return
	}

	SendOk(c, "Article successfully added!")
}

func DeleteArticleHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	deleteCount, err := deleteDocument(articlesLocation, gin.H{"id_name": id})
	if err != nil {
		SendBadRequest(c, "Could not delete document into DB.")
		return
	}

	if deleteCount != 0 {
		SendOk(c, fmt.Sprintf("%d articles were successfully deleted!", deleteCount))
	} else {
		SendOk(c, "No articles were deleted.")
	}
}

func EditArticleHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	var articleUpdate DocumentUpdate
	articleUpdate.Filter = gin.H{"id_name": id}
	c.BindJSON(&articleUpdate.Update)

	editCount, err := editDocument(articlesLocation, articleUpdate)
	if err != nil {
		SendBadRequest(c, err.Error())
		return
	}

	if editCount != 0 {
		SendOk(c, fmt.Sprintf("%d articles were successfully edited!", editCount))
	} else {
		SendOk(c, "No articles were edited.")
	}
}
