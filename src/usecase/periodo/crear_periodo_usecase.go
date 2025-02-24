package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type CrearPeriodoUseCase struct {
	repo domain.PeriodoRepository
}

func NewCrearPeriodoUseCase(repo domain.PeriodoRepository) *CrearPeriodoUseCase {
	return &CrearPeriodoUseCase{repo: repo}
}

func (u *CrearPeriodoUseCase) Execute(request dto.PeriodoDTO) shared.APIResponse {

	existingPeriodo, err := u.repo.FindByNombre(request.Nombre)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el periodo", nil)
	}

	if existingPeriodo.Existe() {
		return shared.NewAPIResponse(400, "Ya existe un periodo con este nombre", nil)
	}

	periodo := domain.NewPeriodo(u.repo)
	periodo.SetNombre(request.Nombre)
	periodo.SetFechaInicio(request.FechaInicio)
	periodo.SetFechaFin(request.FechaFin)

	if err := u.repo.Save(periodo); err != nil {
		return shared.NewAPIResponse(500, "Error al guardar el periodo", nil)
	}

	return shared.NewAPIResponse(201, "Periodo creado exitosamente", periodo.ToDTO())
}
