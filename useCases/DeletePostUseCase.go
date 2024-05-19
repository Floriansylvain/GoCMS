package useCases

import (
	"GoCMS/adapters/secondary/gateways"
	"gorm.io/gorm"
)

type DeletePostUseCase struct {
	postRepository gateways.PostRepository
}

func NewDeletePostUseCase(db *gorm.DB) *DeletePostUseCase {
	return &DeletePostUseCase{
		postRepository: *gateways.NewPostRepository(db),
	}
}

func (g *DeletePostUseCase) DeletePost(userId uint32) error {
	return g.postRepository.Delete(userId)
}
