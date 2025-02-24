package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type ActualizarPeriodoUseCase struct {
	repo domain.PeriodoRepository
}

func NewActualizarPeriodoUseCase(repo domain.PeriodoRepository) *ActualizarPeriodoUseCase {
	return &ActualizarPeriodoUseCase{repo: repo}
}

func (u *ActualizarPeriodoUseCase) Execute(id int64, request dto.PeriodoDTO) shared.APIResponse {

	periodo, err := u.repo.FindByID(id)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el periodo", nil)
	}

	if !periodo.Existe() {
		return shared.NewAPIResponse(404, "El periodo no existe", nil)
	}

	periodo.SetNombre(request.Nombre)
	periodo.SetFechaInicio(request.FechaInicio)
	periodo.SetFechaFin(request.FechaFin)

	if err := u.repo.Update(periodo); err != nil {
		return shared.NewAPIResponse(500, "Error al actualizar el periodo", nil)
	}

	return shared.NewAPIResponse(200, "Periodo actualizado exitosamente", periodo.ToDTO())
}
