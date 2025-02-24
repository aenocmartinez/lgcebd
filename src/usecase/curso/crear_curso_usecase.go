package curso

import (
	"net/http"

	"ebd/src/domain"
	"ebd/src/infraestructure/di"
	"ebd/src/shared"
	formrequest "ebd/src/view/formrequest/curso"
)

type CrearCursoUseCase struct{}

func (uc *CrearCursoUseCase) Execute(request formrequest.CursoFormRequest) shared.APIResponse {
	if err := request.Validate(); err != nil {
		return shared.NewAPIResponse(http.StatusBadRequest, err.Error(), nil)
	}

	cursoRepo := di.GetContainer().GetCursoRepository()

	existingCurso, err := cursoRepo.FindByNombre(request.Nombre)
	if err != nil {
		return shared.NewAPIResponse(http.StatusInternalServerError, "Error al verificar existencia del curso", nil)
	}

	if existingCurso.Existe() {
		return shared.NewAPIResponse(http.StatusConflict, "El curso ya existe", nil)
	}

	curso := domain.NewCurso(cursoRepo)
	curso.SetNombre(request.Nombre)
	curso.SetEdadMinima(request.EdadMinima)
	curso.SetEdadMaxima(request.EdadMaxima)
	curso.SetEstado("activo")

	if err := curso.Save(); err != nil {
		return shared.NewAPIResponse(http.StatusInternalServerError, "Error al guardar el curso", nil)
	}

	return shared.NewAPIResponse(http.StatusCreated, "Curso creado exitosamente", curso.ToDTO())
}
