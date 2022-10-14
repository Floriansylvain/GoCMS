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

var ARTICLES_LOCATION = Location{Database: "gohcms", Collection: "articles"}

// TODO check if id_name is unique among all articles

func AddArticle(c *gin.Context) {
	var article Article
	if c.BindJSON(&article) != nil {
		SendErrorMessageToClient(c, "Could not correctly parse the article.")
	}

	document, err := bson.Marshal(article)
	if err != nil {
		SendErrorMessageToClient(c, "could not correctly marshal the article")
	}

	err = pushDocument(ARTICLES_LOCATION, document)
	if err != nil {
		SendErrorMessageToClient(c, "could not insert document into DB")
	}

	SendOkMessageToClient(c, "Article successfully added!")
}

func GetAllArticles(c *gin.Context) {
	var articles []Article
	documents, err := getDocuments(ARTICLES_LOCATION, bson.D{})
	if err != nil {
		SendErrorMessageToClient(c, err.Error())
	}

	for i := 0; i < len(documents); i++ {
		var newArticle Article
		bson.Unmarshal(documents[i], &newArticle)
		fmt.Println(newArticle)
		articles = append(articles, newArticle)
	}

	c.JSON(http.StatusOK, articles)
}
