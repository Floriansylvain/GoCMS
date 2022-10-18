package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Article struct {
	Id_name string      `json:"id_name" bson:"id_name"`
	Date    int64       `json:"date" bson:"date"`
	Content interface{} `json:"content" bson:"content"`
}

type DelArticle struct {
	Id_name string `json:"id_name" bson:"id_name"`
}

var ARTICLES_LOCATION = Location{Database: "gohcms", Collection: "articles"}

// TODO Change to GetArticleList that will find filter json context
func GetAllArticles(c *gin.Context) {
	var articles []Article
	documents, err := getDocuments(ARTICLES_LOCATION, bson.D{})
	if err != nil {
		SendErrorMessageToClient(c, err.Error())
		return
	}

	for i := 0; i < len(documents); i++ {
		var newArticle Article
		bson.Unmarshal(documents[i], &newArticle)
		articles = append(articles, newArticle)
	}

	c.JSON(http.StatusOK, articles)
}

func IsArticleIdAlreadyUsed(id string) bool {
	documents, _ := getDocuments(ARTICLES_LOCATION, bson.D{})
	for i := 0; i < len(documents); i++ {
		var newArticle Article
		bson.Unmarshal(documents[i], &newArticle)
		if newArticle.Id_name == id {
			return true
		}
	}
	return false
}

func AddArticle(c *gin.Context) {
	var article Article
	if c.BindJSON(&article) != nil {
		SendErrorMessageToClient(c, "Could not correctly parse the article.")
		return
	}

	document, err := bson.Marshal(article)
	if err != nil {
		SendErrorMessageToClient(c, "Could not correctly marshal the article.")
		return
	}

	if IsArticleIdAlreadyUsed(article.Id_name) {
		SendErrorMessageToClient(c, "Article ID already used.")
		return
	}

	err = pushDocument(ARTICLES_LOCATION, document)
	if err != nil {
		SendErrorMessageToClient(c, "Could not insert document into DB.")
		return
	}

	SendOkMessageToClient(c, "Article successfully added!")
}

func DeleteArticle(c *gin.Context) {
	var delArticle DelArticle
	if c.BindJSON(&delArticle) != nil {
		SendErrorMessageToClient(c, "Could not correctly parse the article ID.")
		return
	}

	document, err := bson.Marshal(delArticle)
	if err != nil {
		SendErrorMessageToClient(c, "Could not correctly marshal the article ID.")
		return
	}

	deleteCount, err := deleteDocument(ARTICLES_LOCATION, document)
	if err != nil {
		SendErrorMessageToClient(c, "Could not insert document into DB.")
		return
	}

	SendOkMessageToClient(c, fmt.Sprintf("%d articles were successfully deleted!", deleteCount))
}
