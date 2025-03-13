package useCases

import (
	"GoCMS/adapters/secondary/gateways"
	"GoCMS/domain/post"
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

func (g *GetPostUseCase) GetPostByName(name string) (post.Post, error) {
	return g.postRepository.GetByName(name)
}
