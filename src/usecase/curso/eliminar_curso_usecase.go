package usecase

import (
	"ebd/src/infraestructure/di"
	"ebd/src/shared"
)

type EliminarCursoUseCase struct{}

func (u *EliminarCursoUseCase) Execute(id int64) shared.APIResponse {
	cursoRepo := di.GetContainer().GetCursoRepository()

	curso, err := cursoRepo.FindByID(id)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el curso", nil)
	}

	if !curso.Existe() {
		return shared.NewAPIResponse(404, "Curso no encontrado", nil)
	}

	if err := curso.Delete(); err != nil {
		return shared.NewAPIResponse(500, "Error al eliminar el curso", nil)
	}

	return shared.NewAPIResponse(200, "Curso eliminado exitosamente", nil)
}
