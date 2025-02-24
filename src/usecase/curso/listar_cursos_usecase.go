package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type ListarCursosUseCase struct {
	repo domain.CursoRepository
}

func NewListarCursosUseCase(repo domain.CursoRepository) *ListarCursosUseCase {
	return &ListarCursosUseCase{repo: repo}
}

func (u *ListarCursosUseCase) Execute() shared.APIResponse {

	cursos, err := u.repo.List()
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener los cursos", nil)
	}

	return shared.NewAPIResponse(200, "Cursos obtenidos exitosamente", cursos)
}
