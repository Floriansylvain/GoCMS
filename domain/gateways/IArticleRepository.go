package gateways

import (
	. "GohCMS2/domain/article"
)

type IArticleRepository interface {
	Get(id int) Article
	GetAll() []Article
	Create(article Article) Article
}
