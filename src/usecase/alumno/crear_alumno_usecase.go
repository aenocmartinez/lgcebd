package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	formrequest "ebd/src/view/formrequest/alumno"
)

type CrearAlumnoUseCase struct {
	alumnoRepo domain.AlumnoRepository
}

func NewCrearAlumnoUseCase(alumnoRepo domain.AlumnoRepository) *CrearAlumnoUseCase {
	return &CrearAlumnoUseCase{alumnoRepo: alumnoRepo}
}

func (u *CrearAlumnoUseCase) Execute(request formrequest.AlumnoFormRequest) shared.APIResponse {

	existingAlumno, err := u.alumnoRepo.FindByNombre(request.Nombre)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al verificar el alumno", nil)
	}
	if existingAlumno.Existe() {
		return shared.NewAPIResponse(400, "El alumno ya existe", nil)
	}

	newAlumno := domain.NewAlumno(u.alumnoRepo)
	newAlumno.SetNombre(request.Nombre)
	newAlumno.SetFechaNacimiento(request.FechaNacimiento)
	newAlumno.SetTelefono(request.Telefono)
	newAlumno.SetAcudiente(request.Acudiente)
	newAlumno.SetAcudienteTelefono(request.AcudienteTelefono)
	newAlumno.SetDireccion(request.Direccion)

	if err := newAlumno.Save(); err != nil {
		return shared.NewAPIResponse(500, "Error al guardar el alumno", nil)
	}

	return shared.NewAPIResponse(201, "Alumno creado exitosamente", newAlumno.ToDTO())
}
