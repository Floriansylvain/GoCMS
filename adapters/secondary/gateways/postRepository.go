package gateways

import (
	entity "GohCMS2/adapters/secondary/gateways/models"
	"GohCMS2/domain/gateways"
	domain "GohCMS2/domain/post"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db}
}

func mapPostToDomain(post entity.Post) domain.Post {
	return domain.FromDb(post.ID, post.Title, post.Body, post.CreatedAt, post.UpdatedAt)
}

func (a *PostRepository) Get(id uint32) (domain.Post, error) {
	var article entity.Post
	err := a.db.Model(&entity.Post{}).First(&article, id).Error
	if err != nil {
		return domain.Post{}, err
	}

	return mapPostToDomain(article), nil
}

func (a *PostRepository) Create(article domain.Post) (domain.Post, error) {
	creationResult := a.db.Create(&entity.Post{
		Title: article.Title,
		Body:  article.Body,
	})
	if creationResult.Error != nil {
		return domain.Post{}, creationResult.Error
	}

	var createdArticle entity.Post
	creationResult.Scan(&createdArticle)

	return mapPostToDomain(createdArticle), nil
}

func (a *PostRepository) GetAll() []domain.Post {
	var articles []entity.Post
	err := a.db.Model(&entity.Post{}).Find(&articles).Error
	if err != nil {
		return []domain.Post{}
	}

	var domainArticles = make([]domain.Post, 0)
	for _, article := range articles {
		domainArticles = append(domainArticles, mapPostToDomain(article))
	}

	return domainArticles
}

var _ gateways.IArticleRepository = &PostRepository{}
