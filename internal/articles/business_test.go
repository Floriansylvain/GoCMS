package articles

import (
	"log"
	"testing"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var Article1, _ = bson.Marshal(Article{
	IdName:  "test_article_1",
	Date:    1671482492,
	Content: gin.H{},
})

var Article2, _ = bson.Marshal(Article{
	IdName:  "test_article_2",
	Date:    1671899112,
	Content: gin.H{},
})

var BSONConvertedArticles = [][]byte{Article1, Article2}

func TestGetAllArticlesBusiness(t *testing.T) {
	documents := GetAllArticlesBusiness(BSONConvertedArticles)
	article := documents[0]

	if article.IdName != "test_article_1" {
		log.Fatalf(`Excepted "test_article_1" as IdName, found "%v"`, article.IdName)
	} else if article.Date != 1671482492 {
		log.Fatalf(`Excepted 1671482492 as Date, found "%v"`, article.Date)
	}
}

func TestIsArticleIdAlreadyUsed(t *testing.T) {
	if IsArticleIdAlreadyUsed("test_article_1", BSONConvertedArticles) == false {
		log.Fatalf(`Excepted true, found false`)
	} else if IsArticleIdAlreadyUsed("test_article_3", BSONConvertedArticles) == true {
		log.Fatalf(`Excepted false, found true`)
	}
}
