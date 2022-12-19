package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllArticlesHandler(c *gin.Context) {
	documents, err := getDocuments(articlesLocation, bson.D{})
	if err != nil {
		SendBadRequest(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetAllArticlesBusiness(documents))
}

func GetArticleHandler(c *gin.Context) {
	articleID := c.Params.ByName("id")
	articles, err := getDocuments(articlesLocation,
		bson.D{{Key: "id_name", Value: articleID}})
	if err != nil {
		SendBadRequest(c, err.Error())
		return
	}
	if len(articles) == 0 {
		SendBadRequest(c, "The ID provided doesn't match any article.")
		return
	}
	var parsedArticle Article
	bson.Unmarshal(articles[0], &parsedArticle)
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
		SendBadRequest(c, "Could not insert document into DB.")
		return
	}

	SendOk(c, "Article successfully added!")
}

func DeleteArticleHandler(c *gin.Context) {
	var delArticle DelArticle
	delArticle.IdName = c.Params.ByName("id")
	if c.BindJSON(&delArticle) != nil {
		SendBadRequest(c, "Could not correctly parse the article ID.")
		return
	}

	document, err := bson.Marshal(delArticle)
	if err != nil {
		SendBadRequest(c, "Could not correctly marshal the article ID.")
		return
	}

	deleteCount, err := deleteDocument(articlesLocation, document)
	if err != nil {
		SendBadRequest(c, "Could not insert document into DB.")
		return
	}

	SendOk(c, fmt.Sprintf("%d articles were successfully deleted!", deleteCount))
}
