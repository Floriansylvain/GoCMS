package useCases

import (
	"GohCMS2/adapters/secondary/gateways"
	"GohCMS2/domain/article"
	"gorm.io/gorm"
)

type GetArticleUseCase struct {
	articleRepository gateways.ArticleRepository
}

func NewGetArticleUseCase(db *gorm.DB) *GetArticleUseCase {
	return &GetArticleUseCase{
		articleRepository: *gateways.NewArticleRepository(db),
	}
}

func (g *GetArticleUseCase) GetArticle(id uint32) (article.Article, error) {
	return g.articleRepository.Get(id)
}
