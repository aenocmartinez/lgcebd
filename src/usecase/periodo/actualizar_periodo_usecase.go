package usecase

import (
	"ebd/src/infraestructure/di"
	"ebd/src/shared"

	formrequest "ebd/src/view/formrequest/periodo"
)

type ActualizarPeriodoUseCase struct{}

func (u *ActualizarPeriodoUseCase) Execute(id int64, request formrequest.PeriodoFormRequest) shared.APIResponse {
	periodoRepo := di.GetContainer().GetPeriodoRepository()

	periodo, err := periodoRepo.FindByID(id)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el periodo", nil)
	}
	if !periodo.Existe() {
		return shared.NewAPIResponse(404, "Periodo no encontrado", nil)
	}

	periodo.SetNombre(request.Nombre)
	periodo.SetFechaInicio(request.FechaInicio)
	periodo.SetFechaFin(request.FechaFin)

	if err := periodo.Update(); err != nil {
		return shared.NewAPIResponse(500, "Error al actualizar el periodo", nil)
	}

	return shared.NewAPIResponse(200, "Periodo actualizado exitosamente", periodo.ToDTO())
}
