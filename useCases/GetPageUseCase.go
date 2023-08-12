package useCases

import (
	. "GohCMS2/adapters/secondary/gateways"
)

type GetPageUseCase struct {
	pageRepository PageRepository
}

func NewGetPageUseCase() *GetPageUseCase {
	return &GetPageUseCase{
		pageRepository: *NewPageRepository(),
	}
}

func (g *GetPageUseCase) GetPage(name string, data interface{}) ([]byte, error) {
	return g.pageRepository.Get(name, data)
}
