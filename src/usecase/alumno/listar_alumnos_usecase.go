package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type ListarAlumnosUseCase struct {
	alumnoRepo domain.AlumnoRepository
}

func NewListarAlumnosUseCase(alumnoRepo domain.AlumnoRepository) *ListarAlumnosUseCase {
	return &ListarAlumnosUseCase{alumnoRepo: alumnoRepo}
}

func (u *ListarAlumnosUseCase) Execute() shared.APIResponse {
	alumnosDto := []dto.AlumnoDTO{}
	alumnos, err := u.alumnoRepo.List()
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener los alumnos", nil)
	}

	for _, alumno := range alumnos {
		alumnosDto = append(alumnosDto, *alumno.ToDTO())
	}

	return shared.NewAPIResponse(200, "Alumnos obtenidos exitosamente", alumnosDto)
}
