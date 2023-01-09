package articles

import (
	"github.com/Floriansylvain/GohCMS/internal/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Article struct {
	TitleID string   `json:"titleID" bson:"titleID"`
	Title   string   `json:"title" bson:"title"`
	Date    int64    `json:"date" bson:"date"`
	Content gin.H    `json:"content" bson:"content"`
	Tags    []string `json:"tags" bson:"tags"`
	Online  bool     `json:"online" bson:"online"`
}

var articlesLocation = database.Location{Database: "gohcms", Collection: "articles"}

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
		if newArticle.TitleID == id {
			return true
		}
	}
	return false
}
