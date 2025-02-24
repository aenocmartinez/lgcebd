package curso

import (
	"net/http"

	"ebd/src/infraestructure/di"
	"ebd/src/shared"
)

type ListarCursosUseCase struct{}

func (uc *ListarCursosUseCase) Execute() shared.APIResponse {

	cursoRepo := di.GetContainer().GetCursoRepository()

	cursos, err := cursoRepo.List()
	if err != nil {
		return shared.NewAPIResponse(http.StatusInternalServerError, "Error al obtener cursos", nil)
	}

	if len(cursos) == 0 {
		return shared.NewAPIResponse(http.StatusOK, "No hay cursos disponibles", cursos)
	}

	return shared.NewAPIResponse(http.StatusOK, "Lista de cursos obtenida correctamente", cursos)
}
