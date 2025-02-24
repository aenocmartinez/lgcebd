package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type ActualizarCursoUseCase struct {
	cursoRepo domain.CursoRepository
}

func NewActualizarCursoUseCase(cursoRepo domain.CursoRepository) *ActualizarCursoUseCase {
	return &ActualizarCursoUseCase{cursoRepo: cursoRepo}
}

func (u *ActualizarCursoUseCase) Execute(id int64, request dto.CursoDTO) shared.APIResponse {

	curso, err := u.cursoRepo.FindByID(id)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el curso", nil)
	}

	if !curso.Existe() {
		return shared.NewAPIResponse(404, "Curso no encontrado", nil)
	}

	existente, err := u.cursoRepo.FindByNombre(request.Nombre)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al validar el curso existente", nil)
	}

	if existente.Existe() && existente.GetID() != id {
		return shared.NewAPIResponse(400, "Ya existe un curso con este nombre", nil)
	}

	cursos, err := u.cursoRepo.List()
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener los cursos existentes", nil)
	}

	for _, c := range cursos {
		if c.ID != id &&
			((request.EdadMinima >= c.EdadMinima && request.EdadMinima <= c.EdadMaxima) ||
				(request.EdadMaxima >= c.EdadMinima && request.EdadMaxima <= c.EdadMaxima) ||
				(c.EdadMinima >= request.EdadMinima && c.EdadMinima <= request.EdadMaxima) ||
				(c.EdadMaxima >= request.EdadMinima && c.EdadMaxima <= request.EdadMaxima)) {
			return shared.NewAPIResponse(400, "El curso tiene edades que se cruzan con otro curso existente", nil)
		}
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
