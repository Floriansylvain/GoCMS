package useCases

import (
	"GoCMS/domain/gateways"
	"GoCMS/domain/post"
)

type UpdatePostUseCase struct {
	postRepository gateways.IPostRepository
}

func NewUpdatePostUseCase(postRepository gateways.IPostRepository) *UpdatePostUseCase {
	return &UpdatePostUseCase{postRepository}
}

func (g *UpdatePostUseCase) UpdateBody(id uint32, body string) (post.Post, error) {
	return g.postRepository.UpdateBody(id, body)
}

func (g *UpdatePostUseCase) AddImage(postId uint32, imageId uint32) error {
	return g.postRepository.AddImage(postId, imageId)
}

func (g *UpdatePostUseCase) UpdateIsOnline(id uint32, isOnline bool) (post.Post, error) {
	return g.postRepository.UpdateIsOnline(id, isOnline)
}
