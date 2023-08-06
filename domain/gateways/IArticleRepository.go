package gateways

import (
	. "GohCMS2/domain/article"
)

type IArticleRepository interface {
	Get(id uint32) (Article, error)
	GetAll() []Article
	Create(article Article) (Article, error)
}
