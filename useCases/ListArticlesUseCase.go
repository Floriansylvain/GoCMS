package useCases

import (
	"GohCMS2/adapters/secondary/gateways"
	"GohCMS2/domain/article"
	"gorm.io/gorm"
)

type ListArticlesUseCase struct {
	articleRepository gateways.ArticleRepository
}

func NewListArticlesUseCase(db *gorm.DB) *ListArticlesUseCase {
	return &ListArticlesUseCase{
		articleRepository: *gateways.NewArticleRepository(db),
	}
}

func (g *ListArticlesUseCase) ListArticles() []article.Article {
	return g.articleRepository.GetAll()
}
