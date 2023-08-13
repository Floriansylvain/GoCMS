package gateways

import (
	"GohCMS2/domain/article"
)

type IArticleRepository interface {
	Get(id uint32) (article.Article, error)
	GetAll() []article.Article
	Create(article article.Article) (article.Article, error)
}
