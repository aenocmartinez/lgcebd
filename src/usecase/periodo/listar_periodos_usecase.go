package usecase

import (
	"ebd/src/infraestructure/di"
	"ebd/src/shared"
)

type ListarPeriodosUseCase struct{}

func (u *ListarPeriodosUseCase) Execute() shared.APIResponse {
	periodoRepo := di.GetContainer().GetPeriodoRepository()

	periodos, err := periodoRepo.List()
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener los periodos", nil)
	}

	return shared.NewAPIResponse(200, "Periodos obtenidos exitosamente", periodos)
}
