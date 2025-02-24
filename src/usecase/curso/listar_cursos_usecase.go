package usecase

import (
	"ebd/src/infraestructure/di"
	"ebd/src/shared"
)

type ListarCursosUseCase struct{}

func (u *ListarCursosUseCase) Execute() shared.APIResponse {
	cursoRepo := di.GetContainer().GetCursoRepository()

	cursos, err := cursoRepo.List()
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener los cursos", nil)
	}

	return shared.NewAPIResponse(200, "Cursos obtenidos exitosamente", cursos)
}
