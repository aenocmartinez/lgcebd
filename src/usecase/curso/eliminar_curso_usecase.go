package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type EliminarCursoUseCase struct {
	repo domain.CursoRepository
}

func NewEliminarCursoUseCase(repo domain.CursoRepository) *EliminarCursoUseCase {
	return &EliminarCursoUseCase{repo: repo}
}

func (u *EliminarCursoUseCase) Execute(id int64) shared.APIResponse {
	curso, err := u.repo.FindByID(id)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el curso", nil)
	}

	if !curso.Existe() {
		return shared.NewAPIResponse(404, "El curso no existe", nil)
	}

	if err := u.repo.Delete(id); err != nil {
		return shared.NewAPIResponse(500, "Error al eliminar el curso", nil)
	}

	return shared.NewAPIResponse(200, "Curso eliminado exitosamente", nil)
}
