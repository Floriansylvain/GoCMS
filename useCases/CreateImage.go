package useCases

import (
	"GoCMS/domain/gateways"
	"GoCMS/domain/image"
	"mime/multipart"
)

type CreateImageUseCase struct {
	imageRepository gateways.IImageRepository
}

func NewCreateImageUseCase(imageRepository gateways.IImageRepository) *CreateImageUseCase {
	return &CreateImageUseCase{imageRepository}
}

func (g *CreateImageUseCase) CreateImage(file multipart.File, fileHeader multipart.FileHeader) (image.Image, error) {
	return g.imageRepository.Create(file, fileHeader)
}
