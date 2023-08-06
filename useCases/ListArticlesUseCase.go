package useCases

import (
	. "GohCMS2/adapters/secondary/gateways"
	. "GohCMS2/domain/article"
)

type ListArticlesUseCase struct {
	articleRepository ArticleRepository
}

func NewListArticlesUseCase() *ListArticlesUseCase {
	return &ListArticlesUseCase{
		articleRepository: *NewArticleRepository(),
	}
}

func (g *ListArticlesUseCase) ListArticles() []Article {
	return g.articleRepository.GetAll()
}
