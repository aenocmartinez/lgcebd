package usecase

import (
	"ebd/src/domain"
	"ebd/src/infraestructure/di"
	"ebd/src/shared"
	formrequest "ebd/src/view/formrequest/curso"
)

type CrearCursoUseCase struct{}

func (u *CrearCursoUseCase) Execute(request formrequest.CursoFormRequest) shared.APIResponse {
	cursoRepo := di.GetContainer().GetCursoRepository()

	existente, err := cursoRepo.FindByNombre(request.Nombre)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el curso", nil)
	}
	if existente.Existe() {
		return shared.NewAPIResponse(400, "El curso ya existe", nil)
	}

	curso := domain.NewCurso(cursoRepo)
	curso.SetNombre(request.Nombre)
	curso.SetEdadMinima(request.EdadMinima)
	curso.SetEdadMaxima(request.EdadMaxima)
	curso.SetEstado("activo")

	if err := curso.Save(); err != nil {
		return shared.NewAPIResponse(500, "Error al crear el curso", nil)
	}

	return shared.NewAPIResponse(201, "Curso creado exitosamente", curso.ToDTO())
}
