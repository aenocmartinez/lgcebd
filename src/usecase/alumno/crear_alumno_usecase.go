package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	formrequest "ebd/src/view/formrequest/alumno"
	"log"
)

type CrearAlumnoUseCase struct {
	alumnoRepo       domain.AlumnoRepository
	cursoPeriodoRepo domain.CursoPeriodoRepository
	matriculaRepo    domain.MatriculaRepository
}

func NewCrearAlumnoUseCase(
	alumnoRepo domain.AlumnoRepository,
	cursoPeriodoRepo domain.CursoPeriodoRepository,
	matriculaRepo domain.MatriculaRepository,
) *CrearAlumnoUseCase {
	return &CrearAlumnoUseCase{
		alumnoRepo:       alumnoRepo,
		cursoPeriodoRepo: cursoPeriodoRepo,
		matriculaRepo:    matriculaRepo,
	}
}

func (u *CrearAlumnoUseCase) Execute(request formrequest.AlumnoFormRequest) shared.APIResponse {

	newAlumno := domain.NewAlumno(u.alumnoRepo)
	newAlumno.SetNombre(request.Nombre)
	newAlumno.SetFechaNacimiento(request.FechaNacimiento + "T00:00:00-05:00")
	newAlumno.SetTelefono(request.Telefono)
	newAlumno.SetAcudiente(request.Acudiente)
	newAlumno.SetAcudienteTelefono(request.AcudienteTelefono)
	newAlumno.SetDireccion(request.Direccion)
	newAlumno.SetActivo(true)

	if err := newAlumno.Save(); err != nil {
		return shared.NewAPIResponse(500, "Error al guardar el alumno", nil)
	}

	cursoPeriodoID, err := u.cursoPeriodoRepo.ObtenerPeriodoCursoIDPorEdad(newAlumno.CalcularEdad())
	if err != nil {
		log.Println("Error al obtener el curso periodo por edad:", err)
		return shared.NewAPIResponse(200, "Estudiante creado con éxito pero no se pudo matricular", nil)
	}

	matricularUseCase := NewMatricularAlumnoUseCase(u.alumnoRepo, u.cursoPeriodoRepo, u.matriculaRepo)
	response := matricularUseCase.Execute(newAlumno.GetID(), cursoPeriodoID)

	if response.StatusCode != 200 {
		log.Println("Error al matricular al estudiante:", err)
		return shared.NewAPIResponse(200, "Estudiante creado con éxito pero no se pudo matricular", nil)
	}

	return shared.NewAPIResponse(201, "Alumno creado exitosamente", newAlumno.ToDTO())
}
