package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type ListarPeriodosUseCase struct {
	repo domain.PeriodoRepository
}

func NewListarPeriodosUseCase(repo domain.PeriodoRepository) *ListarPeriodosUseCase {
	return &ListarPeriodosUseCase{repo: repo}
}

func (u *ListarPeriodosUseCase) Execute() shared.APIResponse {
	periodos, err := u.repo.List()
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener los periodos", nil)
	}
	return shared.NewAPIResponse(200, "Periodos obtenidos exitosamente", periodos)
}
