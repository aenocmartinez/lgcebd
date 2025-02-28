package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type ListarCelebracionesUseCase struct {
	celebracionRepo domain.CelebracionRepository
}

func NewListarCelebracionUseCase(repository domain.CelebracionRepository) *ListarCelebracionesUseCase {
	return &ListarCelebracionesUseCase{
		celebracionRepo: repository,
	}
}

func (uc *ListarCelebracionesUseCase) Execute() shared.APIResponse {
	celebraciones := uc.celebracionRepo.List()

	lista := []dto.CelebracionDto{}
	for _, celebracion := range celebraciones {
		lista = append(lista, celebracion.ToDTO())
	}

	return shared.NewAPIResponse(200, "Celebraciones obtenidas exitosamente", lista)
}
