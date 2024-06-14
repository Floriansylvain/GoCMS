package useCases

import (
	"GoCMS/adapters/secondary/gateways"
	"GoCMS/domain/post"
	"gorm.io/gorm"
)

type UpdatePostUseCase struct {
	postRepository gateways.PostRepository
}

func NewUpdatePostUseCase(db *gorm.DB) *UpdatePostUseCase {
	return &UpdatePostUseCase{
		postRepository: *gateways.NewPostRepository(db),
	}
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
