package useCases

import (
	"GohCMS2/adapters/secondary/gateways"
	"GohCMS2/domain/post"
	"gorm.io/gorm"
)

type GetPostUseCase struct {
	postRepository gateways.PostRepository
}

func NewGetPostUseCase(db *gorm.DB) *GetPostUseCase {
	return &GetPostUseCase{
		postRepository: *gateways.NewPostRepository(db),
	}
}

func (g *GetPostUseCase) GetPost(id uint32) (post.Post, error) {
	return g.postRepository.Get(id)
}
