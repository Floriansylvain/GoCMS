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
	var post entity.Post
	err := a.db.Model(&entity.Post{}).First(&post, id).Error
	if err != nil {
		return domain.Post{}, err
	}

	return mapPostToDomain(post), nil
}

func (a *PostRepository) Create(post domain.Post) (domain.Post, error) {
	creationResult := a.db.Create(&entity.Post{
		Title: post.Title,
		Body:  post.Body,
	})
	if creationResult.Error != nil {
		return domain.Post{}, creationResult.Error
	}

	var createdPost entity.Post
	creationResult.Scan(&createdPost)

	return mapPostToDomain(createdPost), nil
}

func (a *PostRepository) GetAll() []domain.Post {
	var posts []entity.Post
	err := a.db.Model(&entity.Post{}).Find(&posts).Error
	if err != nil {
		return []domain.Post{}
	}

	var domainPosts = make([]domain.Post, 0)
	for _, post := range posts {
		domainPosts = append(domainPosts, mapPostToDomain(post))
	}

	return domainPosts
}

var _ gateways.IPostRepository = &PostRepository{}
