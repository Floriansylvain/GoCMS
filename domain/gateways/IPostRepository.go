package gateways

import (
	"GohCMS2/domain/post"
)

type IPostRepository interface {
	Get(id uint32) (post.Post, error)
	GetAll() []post.Post
	Create(post post.Post) (post.Post, error)
}
