package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type ActualizarCursoUseCase struct {
	repo domain.CursoRepository
}

func NewActualizarCursoUseCase(repo domain.CursoRepository) *ActualizarCursoUseCase {
	return &ActualizarCursoUseCase{repo: repo}
}

func (u *ActualizarCursoUseCase) Execute(id int64, request dto.CursoDTO) shared.APIResponse {
	curso, err := u.repo.FindByID(id)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el curso", nil)
	}
	if !curso.Existe() {
		return shared.NewAPIResponse(404, "El curso no existe", nil)
	}

	curso.SetNombre(request.Nombre)
	curso.SetEdadMinima(request.EdadMinima)
	curso.SetEdadMaxima(request.EdadMaxima)
	curso.SetEstado(request.Estado)

	if err := u.repo.Update(curso); err != nil {
		return shared.NewAPIResponse(500, "Error al actualizar el curso", nil)
	}

	return shared.NewAPIResponse(200, "Curso actualizado exitosamente", curso.ToDTO())
}
