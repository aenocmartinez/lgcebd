package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type CrearCursoUseCase struct {
	repo domain.CursoRepository
}

func NewCrearCursoUseCase(repo domain.CursoRepository) *CrearCursoUseCase {
	return &CrearCursoUseCase{repo: repo}
}

func (u *CrearCursoUseCase) Execute(request dto.CursoDTO) shared.APIResponse {
	existingCurso, err := u.repo.FindByNombre(request.Nombre)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el curso", nil)
	}

	if existingCurso.Existe() {
		return shared.NewAPIResponse(400, "Ya existe un curso con este nombre", nil)
	}

	curso := domain.NewCurso(u.repo)
	curso.SetNombre(request.Nombre)
	curso.SetEdadMinima(request.EdadMinima)
	curso.SetEdadMaxima(request.EdadMaxima)
	curso.SetEstado("activo")

	if err := u.repo.Save(curso); err != nil {
		return shared.NewAPIResponse(500, "Error al guardar el curso", nil)
	}

	return shared.NewAPIResponse(201, "Curso creado exitosamente", curso.ToDTO())
}
