package useCases

import (
	"GohCMS2/adapters/secondary/gateways"
	"GohCMS2/domain/post"
	"gorm.io/gorm"
)

type CreatePostUseCase struct {
	postRepository gateways.PostRepository
}

type CreatePostCommand struct {
	Title string
	Body  string
}

func NewCreatePostUseCase(db *gorm.DB) *CreatePostUseCase {
	return &CreatePostUseCase{
		postRepository: *gateways.NewPostRepository(db),
	}
}

func (g *CreatePostUseCase) CreatePost(createPost CreatePostCommand) (post.Post, error) {
	return g.postRepository.Create(post.FromApi(
		createPost.Title,
		createPost.Body,
	))
}
