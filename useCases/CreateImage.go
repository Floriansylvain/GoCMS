package useCases

import (
	"GoCMS/adapters/secondary/gateways"
	"GoCMS/domain/image"
	"gorm.io/gorm"
	"mime/multipart"
)

type CreateImageUseCase struct {
	imageRepository gateways.ImageRepository
}

func NewCreateImageUseCase(db *gorm.DB) *CreateImageUseCase {
	return &CreateImageUseCase{
		imageRepository: *gateways.NewImageRepository(db),
	}
}

func (g *CreateImageUseCase) CreateImage(file multipart.File, fileHeader multipart.FileHeader) (image.Image, error) {
	return g.imageRepository.Create(file, fileHeader)
}
