package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type BuscarAlumnoUseCase struct {
	alumnoRepo domain.AlumnoRepository
}

func NewBuscarAlumnoUseCase(alumnoRepo domain.AlumnoRepository) *BuscarAlumnoUseCase {
	return &BuscarAlumnoUseCase{alumnoRepo: alumnoRepo}
}

func (u *BuscarAlumnoUseCase) Execute(id int64) shared.APIResponse {

	existingAlumno, err := u.alumnoRepo.FindByID(id)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el alumno", nil)
	}

	if !existingAlumno.Existe() {
		return shared.NewAPIResponse(404, "Alumno no encontrado", nil)
	}

	return shared.NewAPIResponse(200, "Alumno encontrado", existingAlumno.ToDTO())
}
