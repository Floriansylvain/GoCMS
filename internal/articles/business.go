package articles

import (
	"github.com/Floriansylvain/GohCMS/internal/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Article struct {
	IdName  string `json:"id_name" bson:"id_name"`
	Date    int64  `json:"date" bson:"date"`
	Content gin.H  `json:"content" bson:"content"`
	PageID  string `json:"page_id" bson:"page_id"`
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
		if newArticle.IdName == id {
			return true
		}
	}
	return false
}
