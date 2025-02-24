package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type AgregarCursoAPeriodoUseCase struct {
	periodoRepo domain.PeriodoRepository
}

func NewAgregarCursoAPeriodoUseCase(periodoRepo domain.PeriodoRepository) *AgregarCursoAPeriodoUseCase {
	return &AgregarCursoAPeriodoUseCase{periodoRepo: periodoRepo}
}

func (u *AgregarCursoAPeriodoUseCase) Execute(periodoID, cursoID int64) shared.APIResponse {
	err := u.periodoRepo.AgregarCurso(periodoID, cursoID)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al asociar el curso al periodo", nil)
	}

	return shared.NewAPIResponse(200, "Curso asociado al periodo exitosamente", nil)
}
