package useCases

import (
	"GoCMS/domain/gateways"
	"GoCMS/domain/post"
)

type GetPostUseCase struct {
	postRepository gateways.IPostRepository
}

func NewGetPostUseCase(postRepository gateways.IPostRepository) *GetPostUseCase {
	return &GetPostUseCase{postRepository}
}

func (g *GetPostUseCase) GetPost(id uint32) (post.Post, error) {
	return g.postRepository.Get(id)
}

func (g *GetPostUseCase) GetPostByName(name string) (post.Post, error) {
	return g.postRepository.GetByName(name)
}
