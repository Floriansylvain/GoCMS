package internal

import (
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

func GetAllArticlesBusiness(documents [][]byte) []Article {
	var articles = []Article{}

	for i := 0; i < len(documents); i++ {
		var newArticle Article
		bson.Unmarshal(documents[i], &newArticle)
		articles = append(articles, newArticle)
	}

	return articles
}

func IsArticleIdAlreadyUsed(id string, documents [][]byte) bool {
	for i := 0; i < len(documents); i++ {
		var newArticle Article
		bson.Unmarshal(documents[i], &newArticle)
		if newArticle.IdName == id {
			return true
		}
	}
	return false
}
