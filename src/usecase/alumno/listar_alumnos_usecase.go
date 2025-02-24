package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type ListarAlumnosUseCase struct {
	alumnoRepo domain.AlumnoRepository
}

func NewListarAlumnosUseCase(alumnoRepo domain.AlumnoRepository) *ListarAlumnosUseCase {
	return &ListarAlumnosUseCase{alumnoRepo: alumnoRepo}
}

func (u *ListarAlumnosUseCase) Execute() shared.APIResponse {
	alumnos, err := u.alumnoRepo.List()
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener los alumnos", nil)
	}

	return shared.NewAPIResponse(200, "Alumnos obtenidos exitosamente", alumnos)
}
