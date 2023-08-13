package useCases

import (
	"GohCMS2/adapters/secondary/gateways"
	"GohCMS2/domain/article"
	"gorm.io/gorm"
)

type CreateArticleUseCase struct {
	articleRepository gateways.ArticleRepository
}

type CreateArticleCommand struct {
	Title string
	Body  string
}

func NewCreateArticleUseCase(db *gorm.DB) *CreateArticleUseCase {
	return &CreateArticleUseCase{
		articleRepository: *gateways.NewArticleRepository(db),
	}
}

func (g *CreateArticleUseCase) CreateArticle(createArticle CreateArticleCommand) (article.Article, error) {
	return g.articleRepository.Create(article.FromApi(
		createArticle.Title,
		createArticle.Body,
	))
}
