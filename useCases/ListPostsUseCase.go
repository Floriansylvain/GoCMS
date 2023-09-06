package useCases

import (
	"GohCMS2/adapters/secondary/gateways"
	"GohCMS2/domain/post"
	"gorm.io/gorm"
)

type ListPostsUseCase struct {
	postRepository gateways.PostRepository
}

func NewListPostsUseCase(db *gorm.DB) *ListPostsUseCase {
	return &ListPostsUseCase{
		postRepository: *gateways.NewPostRepository(db),
	}
}

func (g *ListPostsUseCase) ListPosts() []post.Post {
	return g.postRepository.GetAll()
}
