package gateways

import (
	entity "GoCMS/adapters/secondary/gateways/models"
	"GoCMS/domain/gateways"
	domain "GoCMS/domain/post"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db}
}

func mapPostToDomain(post entity.Post) domain.Post {
	return domain.FromDb(post.ID, post.Title, post.Body, post.Images, post.IsOnline, post.CreatedAt, post.UpdatedAt)
}

func (a *PostRepository) Get(id uint32) (domain.Post, error) {
	var post entity.Post
	err := a.db.Model(&entity.Post{}).Preload("Images").First(&post, id).Error
	if err != nil {
		return domain.Post{}, err
	}

	return mapPostToDomain(post), nil
}

func (a *PostRepository) GetByName(name string) (domain.Post, error) {
	var post entity.Post
	err := a.db.Model(&entity.Post{}).Where("title = ?", name).First(&post).Error
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

func (a *PostRepository) UpdateBody(id uint32, body string) (domain.Post, error) {
	var localPost entity.Post
	err := a.db.Model(&entity.Post{}).First(&localPost, id).Error
	if err != nil {
		return domain.Post{}, err
	}

	localPost.Body = body
	err = a.db.Save(&localPost).Error
	if err != nil {
		return domain.Post{}, err
	}

	newPost := domain.FromDb(
		localPost.ID,
		localPost.Title,
		localPost.Body,
		localPost.Images,
		localPost.IsOnline,
		localPost.CreatedAt,
		localPost.UpdatedAt,
	)

	return newPost, nil
}

func (a *PostRepository) UpdateIsOnline(id uint32, isOnline bool) (domain.Post, error) {
	var localPost entity.Post
	err := a.db.Model(&entity.Post{}).First(&localPost, id).Error
	if err != nil {
		return domain.Post{}, err
	}

	localPost.IsOnline = isOnline
	err = a.db.Save(&localPost).Error
	if err != nil {
		return domain.Post{}, err
	}

	return mapPostToDomain(localPost), nil
}

func (a *PostRepository) Delete(id uint32) error {
	return a.db.Delete(&entity.Post{}, id).Error
}

func (a *PostRepository) AddImage(postId uint32, imageId uint32) error {
	var localPost entity.Post
	err := a.db.Model(&entity.Post{}).First(&localPost, postId).Error
	if err != nil {
		return err
	}

	var localImage entity.Image
	err = a.db.Model(&entity.Image{}).First(&localImage, imageId).Error
	if err != nil {
		return err
	}

	err = a.db.Model(&localPost).Association("Images").Append(&localImage)
	if err != nil {
		return err
	}

	return nil
}

var _ gateways.IPostRepository = &PostRepository{}
