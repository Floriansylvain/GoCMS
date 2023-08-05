package useCases

import (
	. "GohCMS2/adapters/secondary/gateways"
	. "GohCMS2/domain/article"
)

type GetArticleUseCase struct {
	articleRepository ArticleRepository
}

func NewGetArticleUseCase() *GetArticleUseCase {
	return &GetArticleUseCase{
		articleRepository: *NewArticleRepository(),
	}
}

func (g *GetArticleUseCase) GetArticle(id int) Article {
	return g.articleRepository.Get(id)
}
