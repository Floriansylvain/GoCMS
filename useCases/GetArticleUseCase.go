package useCases

import (
	. "GohCMS2/adapters/secondary/gateways"
	. "GohCMS2/domain/article"
	"gorm.io/gorm"
)

type GetArticleUseCase struct {
	articleRepository ArticleRepository
}

func NewGetArticleUseCase(db *gorm.DB) *GetArticleUseCase {
	return &GetArticleUseCase{
		articleRepository: *NewArticleRepository(db),
	}
}

func (g *GetArticleUseCase) GetArticle(id uint32) (Article, error) {
	return g.articleRepository.Get(id)
}
