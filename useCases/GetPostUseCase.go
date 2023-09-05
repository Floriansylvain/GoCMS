package useCases

import (
	"GohCMS2/adapters/secondary/gateways"
	"GohCMS2/domain/post"
	"gorm.io/gorm"
)

type GetArticleUseCase struct {
	articleRepository gateways.PostRepository
}

func NewGetArticleUseCase(db *gorm.DB) *GetArticleUseCase {
	return &GetArticleUseCase{
		articleRepository: *gateways.NewPostRepository(db),
	}
}

func (g *GetArticleUseCase) GetArticle(id uint32) (post.Post, error) {
	return g.articleRepository.Get(id)
}
