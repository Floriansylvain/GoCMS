package useCases

import (
	"GoCMS/adapters/secondary/gateways"
)

type GetPageUseCase struct {
	pageRepository gateways.PageRepository
}

func NewGetPageUseCase() *GetPageUseCase {
	return &GetPageUseCase{
		pageRepository: *gateways.NewPageRepository(),
	}
}

func (g *GetPageUseCase) GetPage(name string, data interface{}) ([]byte, error) {
	return g.pageRepository.Get(name, data)
}
