package gateways

import (
	entity "GohCMS2/adapters/secondary/gateways/models"
	domain "GohCMS2/domain/article"
	"GohCMS2/domain/gateways"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db}
}

func mapArticleToDomain(article entity.Article) domain.Article {
	return domain.FromDb(article.ID, article.Title, article.Body, article.CreatedAt, article.UpdatedAt)
}

func (a *ArticleRepository) Get(id uint32) (domain.Article, error) {
	var article entity.Article
	err := a.db.Model(&entity.Article{}).First(&article, id).Error
	if err != nil {
		return domain.Article{}, err
	}

	return mapArticleToDomain(article), nil
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

	return mapArticleToDomain(createdArticle), nil
}

func (a *ArticleRepository) GetAll() []domain.Article {
	var articles []entity.Article
	err := a.db.Model(&entity.Article{}).Find(&articles).Error
	if err != nil {
		return []domain.Article{}
	}

	var domainArticles = make([]domain.Article, 0)
	for _, article := range articles {
		domainArticles = append(domainArticles, mapArticleToDomain(article))
	}

	return domainArticles
}

var _ gateways.IArticleRepository = &ArticleRepository{}
