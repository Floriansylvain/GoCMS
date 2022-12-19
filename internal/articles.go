package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Article struct {
	IdName  string      `json:"id_name" bson:"id_name"`
	Date    int64       `json:"date" bson:"date"`
	Content interface{} `json:"content" bson:"content"`
}

type DelArticle struct {
	IdName string `json:"id_name" bson:"id_name"`
}

var articlesLocation = Location{Database: "gohcms", Collection: "articles"}

func GetAllArticles(c *gin.Context) {
	// TODO Change to GetArticleList that will find filter json context
	var articles []Article
	documents, err := getDocuments(articlesLocation, bson.D{})
	if err != nil {
		SendBadRequest(c, err.Error())
		return
	}

	for i := 0; i < len(documents); i++ {
		var newArticle Article
		bson.Unmarshal(documents[i], &newArticle)
		articles = append(articles, newArticle)
	}

	c.JSON(http.StatusOK, articles)
}

func GetArticle(c *gin.Context) {
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

func IsArticleIdAlreadyUsed(id string) bool {
	documents, _ := getDocuments(articlesLocation, bson.D{})
	for i := 0; i < len(documents); i++ {
		var newArticle Article
		bson.Unmarshal(documents[i], &newArticle)
		if newArticle.IdName == id {
			return true
		}
	}
	return false
}

func AddArticle(c *gin.Context) {
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

	if IsArticleIdAlreadyUsed(article.IdName) {
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

func DeleteArticle(c *gin.Context) {
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
