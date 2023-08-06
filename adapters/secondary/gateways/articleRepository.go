package gateways

import (
	entity "GohCMS2/adapters/secondary/gateways/models"
	domain "GohCMS2/domain/article"
	. "GohCMS2/domain/gateways"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	a := ArticleRepository{db}
	return &a
}

func (a *ArticleRepository) Get(id uint32) (domain.Article, error) {
	var article domain.Article
	err := a.db.Model(&entity.Article{}).First(&article, id).Error
	if err != nil {
		return article, err
	}

	return domain.FromDb(article.ID, article.Title, article.Body, article.CreatedAt, article.UpdatedAt), nil
}

func (a *ArticleRepository) Create(article domain.Article) (domain.Article, error) {
	creationResult := a.db.Create(&entity.Article{
		Title: article.Title,
		Body:  article.Body,
	})
	if creationResult.Error != nil {
		return domain.Article{}, creationResult.Error
	}

	var createdArticle entity.Article
	creationResult.Scan(&createdArticle)

	return domain.FromDb(
			createdArticle.ID,
			createdArticle.Title,
			createdArticle.Body,
			createdArticle.CreatedAt,
			createdArticle.UpdatedAt),
		nil
}

func (a *ArticleRepository) GetAll() []domain.Article {
	var articles []domain.Article
	err := a.db.Model(&entity.Article{}).Find(&articles).Error
	if err != nil {
		return []domain.Article{}
	}

	return articles
}

var _ IArticleRepository = &ArticleRepository{}
