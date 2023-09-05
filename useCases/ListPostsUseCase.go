package useCases

import (
	"GohCMS2/adapters/secondary/gateways"
	"GohCMS2/domain/post"
	"gorm.io/gorm"
)

type ListArticlesUseCase struct {
	articleRepository gateways.PostRepository
}

func NewListArticlesUseCase(db *gorm.DB) *ListArticlesUseCase {
	return &ListArticlesUseCase{
		articleRepository: *gateways.NewPostRepository(db),
	}
}

func (g *ListArticlesUseCase) ListArticles() []post.Post {
	return g.articleRepository.GetAll()
}
