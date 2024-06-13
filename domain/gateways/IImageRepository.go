package gateways

import (
	"GoCMS/domain/image"
	"mime/multipart"
)

type IImageRepository interface {
	Create(file multipart.File, fileHeader multipart.FileHeader) (image.Image, error)
	Delete(id uint32) error
}
