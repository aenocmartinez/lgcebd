package usecase

import (
	"ebd/src/infraestructure/di"
	"ebd/src/shared"
	formrequest "ebd/src/view/formrequest/curso"
)

type UpdateCursoUseCase struct{}

func (u *UpdateCursoUseCase) Execute(id int64, request formrequest.CursoFormRequest) shared.APIResponse {
	cursoRepo := di.GetContainer().GetCursoRepository()

	curso, err := cursoRepo.FindByID(id)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el curso", nil)
	}

	if !curso.Existe() {
		return shared.NewAPIResponse(404, "Curso no encontrado", nil)
	}

	curso.SetNombre(request.Nombre)
	curso.SetEdadMinima(request.EdadMinima)
	curso.SetEdadMaxima(request.EdadMaxima)
	curso.SetEstado(request.Estado)

	if err := curso.Update(); err != nil {
		return shared.NewAPIResponse(500, "Error al actualizar el curso", nil)
	}

	return shared.NewAPIResponse(200, "Curso actualizado exitosamente", curso.ToDTO())
}
