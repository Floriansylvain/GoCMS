package gateways

import (
	"GohCMS2/db"
	. "GohCMS2/domain/article"
	. "GohCMS2/domain/gateways"
	"context"
)

type ArticleRepository struct {
	client *db.PrismaClient
}

func NewArticleRepository() *ArticleRepository {
	a := ArticleRepository{
		client: db.NewClient(),
	}
	return &a
}

func (a *ArticleRepository) Get(id int) Article {
	a.client.Connect()
	defer a.client.Disconnect()

	article, err := a.client.Article.FindUnique(db.Article.ID.Equals(id)).Exec(context.Background())
	if err != nil {
		panic(err)
	}

	return FromDb(article.ID, article.Title, article.Body, article.CreatedAt.String(), article.UpdatedAt.String())
}

func (a *ArticleRepository) Create(article Article) Article {
	a.client.Connect()
	defer a.client.Disconnect()

	articleDb, err := a.client.Article.CreateOne(
		db.Article.Title.Set(article.Title),
		db.Article.Body.Set(article.Body),
	).Exec(context.Background())
	if err != nil {
		panic(err)
	}

	return FromDb(articleDb.ID, articleDb.Title, articleDb.Body, articleDb.CreatedAt.String(), articleDb.UpdatedAt.String())
}

var _ IArticleRepository = &ArticleRepository{}
