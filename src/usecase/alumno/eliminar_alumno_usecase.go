package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type EliminarAlumnoUseCase struct {
	alumnoRepo domain.AlumnoRepository
}

func NewEliminarAlumnoUseCase(alumnoRepo domain.AlumnoRepository) *EliminarAlumnoUseCase {
	return &EliminarAlumnoUseCase{alumnoRepo: alumnoRepo}
}

func (u *EliminarAlumnoUseCase) Execute(id int64) shared.APIResponse {

	existingAlumno, err := u.alumnoRepo.FindByID(id)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el alumno", nil)
	}

	if !existingAlumno.Existe() {
		return shared.NewAPIResponse(404, "Alumno no encontrado", nil)
	}

	if err := u.alumnoRepo.Delete(id); err != nil {
		return shared.NewAPIResponse(500, "Error al eliminar el alumno", nil)
	}

	return shared.NewAPIResponse(200, "Alumno eliminado exitosamente", nil)
}
