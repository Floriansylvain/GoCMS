package useCases

import (
	"GoCMS/domain/gateways"
	"GoCMS/domain/post"
)

type CreatePostUseCase struct {
	postRepository gateways.IPostRepository
}

type CreatePostCommand struct {
	Title string
	Body  string
}

func NewCreatePostUseCase(postRepository gateways.IPostRepository) *CreatePostUseCase {
	return &CreatePostUseCase{postRepository}
}

func (g *CreatePostUseCase) CreatePost(createPost CreatePostCommand) (post.Post, error) {
	return g.postRepository.Create(post.FromApi(
		createPost.Title,
		createPost.Body,
	))
}
