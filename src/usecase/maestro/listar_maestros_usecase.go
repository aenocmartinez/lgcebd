package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type ListarMaestrosUseCase struct {
	maestroRepo domain.MaestroRepository
}

func NewListarMaestrosUseCase(maestroRepo domain.MaestroRepository) *ListarMaestrosUseCase {
	return &ListarMaestrosUseCase{maestroRepo: maestroRepo}
}

func (u *ListarMaestrosUseCase) Execute() shared.APIResponse {
	maestros, err := u.maestroRepo.List()
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener los maestros", nil)
	}

	maestrosDTO := []dto.MaestroDTO{}
	for _, maestro := range maestros {
		maestrosDTO = append(maestrosDTO, *maestro.ToDTO())
	}

	return shared.NewAPIResponse(200, "Maestros obtenidos correctamente", maestrosDTO)
}
