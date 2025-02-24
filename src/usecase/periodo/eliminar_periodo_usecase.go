package usecase

import (
	"ebd/src/infraestructure/di"
	"ebd/src/shared"
)

type EliminarPeriodoUseCase struct{}

func (u *EliminarPeriodoUseCase) Execute(id int64) shared.APIResponse {
	periodoRepo := di.GetContainer().GetPeriodoRepository()

	periodo, err := periodoRepo.FindByID(id)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el periodo", nil)
	}
	if !periodo.Existe() {
		return shared.NewAPIResponse(404, "Periodo no encontrado", nil)
	}

	if err := periodo.Delete(); err != nil {
		return shared.NewAPIResponse(500, "Error al eliminar el periodo", nil)
	}

	return shared.NewAPIResponse(200, "Periodo eliminado exitosamente", nil)
}
