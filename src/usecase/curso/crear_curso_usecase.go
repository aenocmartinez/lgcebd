package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type CrearCursoUseCase struct {
	cursoRepo domain.CursoRepository
}

func NewCrearCursoUseCase(cursoRepo domain.CursoRepository) *CrearCursoUseCase {
	return &CrearCursoUseCase{cursoRepo: cursoRepo}
}

func (u *CrearCursoUseCase) Execute(request dto.CursoDTO) shared.APIResponse {
	// Verificar si ya existe un curso con el mismo nombre
	existente, err := u.cursoRepo.FindByNombre(request.Nombre)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al validar el curso existente", nil)
	}

	if existente.Existe() {
		return shared.NewAPIResponse(400, "Ya existe un curso con este nombre", nil)
	}

	cursos, err := u.cursoRepo.List()
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener los cursos existentes", nil)
	}

	for _, curso := range cursos {
		if (request.EdadMinima >= curso.EdadMinima && request.EdadMinima <= curso.EdadMaxima) ||
			(request.EdadMaxima >= curso.EdadMinima && request.EdadMaxima <= curso.EdadMaxima) ||
			(curso.EdadMinima >= request.EdadMinima && curso.EdadMinima <= request.EdadMaxima) ||
			(curso.EdadMaxima >= request.EdadMinima && curso.EdadMaxima <= request.EdadMaxima) {
			return shared.NewAPIResponse(400, "El curso tiene edades que se cruzan con otro curso existente", nil)
		}
	}

	curso := domain.NewCurso(u.cursoRepo)
	curso.SetNombre(request.Nombre)
	curso.SetEdadMinima(request.EdadMinima)
	curso.SetEdadMaxima(request.EdadMaxima)
	curso.SetEstado("activo")

	if err := curso.Save(); err != nil {
		return shared.NewAPIResponse(500, "Error al guardar el curso", nil)
	}

	return shared.NewAPIResponse(201, "Curso creado exitosamente", curso.ToDTO())
}
