package useCases

import (
	. "GohCMS2/adapters/secondary/gateways"
	. "GohCMS2/domain/article"
	"gorm.io/gorm"
)

type CreateArticleUseCase struct {
	articleRepository ArticleRepository
}

type CreateArticleCommand struct {
	Title string
	Body  string
}

func NewCreateArticleUseCase(db *gorm.DB) *CreateArticleUseCase {
	return &CreateArticleUseCase{
		articleRepository: *NewArticleRepository(db),
	}
}

func (g *CreateArticleUseCase) CreateArticle(article CreateArticleCommand) (Article, error) {
	return g.articleRepository.Create(FromApi(article.Title, article.Body))
}
