package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type CrearPeriodoUseCase struct {
	periodoRepo domain.PeriodoRepository
	cursoRepo   domain.CursoRepository
}

func NewCrearPeriodoUseCase(
	periodoRepo domain.PeriodoRepository,
	cursoRepo domain.CursoRepository,
) *CrearPeriodoUseCase {
	return &CrearPeriodoUseCase{
		periodoRepo: periodoRepo,
		cursoRepo:   cursoRepo,
	}
}

func (u *CrearPeriodoUseCase) Execute(request dto.PeriodoDTO) shared.APIResponse {

	periodo, err := u.periodoRepo.FindByNombre(request.Nombre)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al verificar la existencia del periodo", nil)
	}

	if periodo.Existe() {
		return shared.NewAPIResponse(400, "Ya existe un periodo con ese nombre", nil)
	}

	nuevoPeriodo := domain.NewPeriodo(u.periodoRepo)
	nuevoPeriodo.SetNombre(request.Nombre)
	nuevoPeriodo.SetFechaInicio(request.FechaInicio)
	nuevoPeriodo.SetFechaFin(request.FechaFin)

	err = u.periodoRepo.Save(nuevoPeriodo)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al crear el periodo", nil)
	}

	cursos, err := u.cursoRepo.List()
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener los cursos", nil)
	}

	if len(cursos) == 0 {
		return shared.NewAPIResponse(200, "Periodo creado sin cursos asociados, no hay cursos habilitados.", nil)
	}

	for _, curso := range cursos {

		if curso.Estado == "inactivo" {
			continue
		}

		err := u.periodoRepo.AgregarCurso(nuevoPeriodo.GetID(), curso.ID)
		if err != nil {
			return shared.NewAPIResponse(500, "Error al asociar el curso al periodo", nil)
		}
	}

	return shared.NewAPIResponse(200, "Periodo creado y cursos asociados exitosamente.", nuevoPeriodo.ToDTO())
}
