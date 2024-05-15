package gateways

import (
	"GoCMS/domain/post"
)

type IPostRepository interface {
	Get(id uint32) (post.Post, error)
	GetAll() []post.Post
	Create(post post.Post) (post.Post, error)
}
