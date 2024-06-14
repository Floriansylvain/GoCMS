package gateways

import (
	"GoCMS/domain/post"
)

type IPostRepository interface {
	Get(id uint32) (post.Post, error)
	GetByName(name string) (post.Post, error)
	GetAll() []post.Post
	Create(post post.Post) (post.Post, error)
	UpdateBody(id uint32, body string) (post.Post, error)
	UpdateIsOnline(id uint32, isOnline bool) (post.Post, error)
	Delete(id uint32) error
	AddImage(postId uint32, imageId uint32) error
}
