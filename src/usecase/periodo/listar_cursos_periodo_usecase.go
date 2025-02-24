package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type ListarCursosPeriodoUseCase struct {
	periodoRepo domain.PeriodoRepository
}

func NewListarCursosPeriodoUseCase(periodoRepo domain.PeriodoRepository) *ListarCursosPeriodoUseCase {
	return &ListarCursosPeriodoUseCase{periodoRepo: periodoRepo}
}

func (u *ListarCursosPeriodoUseCase) Execute(periodoID int64) shared.APIResponse {

	periodo, err := u.periodoRepo.FindByID(periodoID)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener el periodo", nil)
	}

	if !periodo.Existe() {
		return shared.NewAPIResponse(404, "El periodo no existe", nil)
	}

	cursos, err := u.periodoRepo.ObtenerCursos(periodoID)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener los cursos del periodo", nil)
	}

	return shared.NewAPIResponse(200, "Cursos del periodo obtenidos correctamente", cursos)
}
