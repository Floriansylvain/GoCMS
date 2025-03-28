package useCases

import (
	"GoCMS/domain/gateways"
	"GoCMS/domain/post"
)

type ListPostsUseCase struct {
	postRepository gateways.IPostRepository
}

func NewListPostsUseCase(postRepository gateways.IPostRepository) *ListPostsUseCase {
	return &ListPostsUseCase{postRepository}
}

func (g *ListPostsUseCase) ListPosts() []post.Post {
	return g.postRepository.GetAll()
}
