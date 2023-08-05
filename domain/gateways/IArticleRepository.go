package gateways

import (
	. "GohCMS2/domain/article"
)

type IArticleRepository interface {
	Get(id int) Article
	Create(article Article) Article
}
