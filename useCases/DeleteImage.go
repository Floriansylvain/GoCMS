package useCases

import (
	"GoCMS/adapters/secondary/gateways"
	"gorm.io/gorm"
)

type DeleteImageUseCase struct {
	imageRepository gateways.ImageRepository
}

func NewDeleteImageUseCase(db *gorm.DB) *DeleteImageUseCase {
	return &DeleteImageUseCase{
		imageRepository: *gateways.NewImageRepository(db),
	}
}

func (g *DeleteImageUseCase) DeleteImage(imageId uint32) error {
	return g.imageRepository.Delete(imageId)
}
