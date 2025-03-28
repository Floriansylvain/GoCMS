package useCases

import (
	"GoCMS/domain/gateways"
)

type DeletePostUseCase struct {
	postRepository gateways.IPostRepository
}

func NewDeletePostUseCase(postRepository gateways.IPostRepository) *DeletePostUseCase {
	return &DeletePostUseCase{postRepository}
}

func (g *DeletePostUseCase) DeletePost(userId uint32) error {
	return g.postRepository.Delete(userId)
}
