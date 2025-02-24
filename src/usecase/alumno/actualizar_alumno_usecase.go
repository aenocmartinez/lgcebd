package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	formrequest "ebd/src/view/formrequest/alumno"
)

type ActualizarAlumnoUseCase struct {
	alumnoRepo domain.AlumnoRepository
}

func NewActualizarAlumnoUseCase(alumnoRepo domain.AlumnoRepository) *ActualizarAlumnoUseCase {
	return &ActualizarAlumnoUseCase{alumnoRepo: alumnoRepo}
}

func (u *ActualizarAlumnoUseCase) Execute(id int64, request formrequest.AlumnoFormRequest) shared.APIResponse {

	existingAlumno, err := u.alumnoRepo.FindByID(id)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el alumno", nil)
	}

	if !existingAlumno.Existe() {
		return shared.NewAPIResponse(404, "Alumno no encontrado", nil)
	}

	existingAlumno.SetNombre(request.Nombre)
	existingAlumno.SetFechaNacimiento(request.FechaNacimiento)
	existingAlumno.SetTelefono(request.Telefono)
	existingAlumno.SetAcudiente(request.Acudiente)
	existingAlumno.SetAcudienteTelefono(request.AcudienteTelefono)
	existingAlumno.SetDireccion(request.Direccion)

	if err := existingAlumno.Update(); err != nil {
		return shared.NewAPIResponse(500, "Error al actualizar el alumno", nil)
	}

	return shared.NewAPIResponse(200, "Alumno actualizado exitosamente", existingAlumno.ToDTO())
}
