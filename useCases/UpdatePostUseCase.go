package useCases

import (
	"GoCMS/adapters/secondary/gateways"
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

func (g *UpdatePostUseCase) UpdateBody(id uint32, body string) error {
	return g.postRepository.UpdateBody(id, body)
}
