package gateways

import (
	"GohCMS2/domain/post"
)

type IArticleRepository interface {
	Get(id uint32) (post.Post, error)
	GetAll() []post.Post
	Create(article post.Post) (post.Post, error)
}
