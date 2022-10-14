package internal

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Article struct {
	Id_name string      `json:"id_name" bson:"id_name"`
	Date    int64       `json:"date" bson:"date"`
	Content interface{} `json:"content" bson:"content"`
}

func AddArticle(c *gin.Context) {
	var article Article
	if c.BindJSON(&article) != nil {
		SendErrorMessageToClient(c, "Could not correctly parse the article.")
	}

	document, err := bson.Marshal(article)
	if err != nil {
		SendErrorMessageToClient(c, "Could not correctly marshal the article.")
	}

	err = pushDocument(Location{Database: "gohcms", Collection: "articles"}, document)
	if err != nil {
		SendErrorMessageToClient(c, "Could not insert document into DB.")
	}
}
