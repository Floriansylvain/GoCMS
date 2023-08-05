package useCases

import (
	. "GohCMS2/adapters/secondary/gateways"
	. "GohCMS2/domain/article"
)

type CreateArticleUseCase struct {
	articleRepository ArticleRepository
}

type CreateArticleCommand struct {
	Title string
	Body  string
}

func NewCreateArticleUseCase() *CreateArticleUseCase {
	return &CreateArticleUseCase{
		articleRepository: *NewArticleRepository(),
	}
}

func (g *CreateArticleUseCase) CreateAarticle(article CreateArticleCommand) Article {
	return g.articleRepository.Create(FromApi(article.Title, article.Body))
}
