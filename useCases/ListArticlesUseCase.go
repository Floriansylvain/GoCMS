package useCases

import (
	. "GohCMS2/adapters/secondary/gateways"
	. "GohCMS2/domain/article"
	"gorm.io/gorm"
)

type ListArticlesUseCase struct {
	articleRepository ArticleRepository
}

func NewListArticlesUseCase(db *gorm.DB) *ListArticlesUseCase {
	return &ListArticlesUseCase{
		articleRepository: *NewArticleRepository(db),
	}
}

func (g *ListArticlesUseCase) ListArticles() []Article {
	return g.articleRepository.GetAll()
}
