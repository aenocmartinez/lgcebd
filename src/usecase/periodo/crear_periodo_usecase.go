package usecase

import (
	"ebd/src/domain"
	"ebd/src/infraestructure/di"
	"ebd/src/shared"

	formrequest "ebd/src/view/formrequest/periodo"
)

type CrearPeriodoUseCase struct{}

func (u *CrearPeriodoUseCase) Execute(request formrequest.PeriodoFormRequest) shared.APIResponse {
	periodoRepo := di.GetContainer().GetPeriodoRepository()

	existingPeriodo, err := periodoRepo.FindByNombre(request.Nombre)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al verificar la existencia del periodo", nil)
	}
	if existingPeriodo.Existe() {
		return shared.NewAPIResponse(400, "El periodo ya existe", nil)
	}

	periodo := domain.NewPeriodo(periodoRepo)
	periodo.SetNombre(request.Nombre)
	periodo.SetFechaInicio(request.FechaInicio)
	periodo.SetFechaFin(request.FechaFin)

	if err := periodo.Save(); err != nil {
		return shared.NewAPIResponse(500, "Error al guardar el periodo", nil)
	}

	return shared.NewAPIResponse(201, "Periodo creado exitosamente", periodo.ToDTO())
}
