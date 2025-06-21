package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
)

type ListarAlumnosMatriculadosUseCase struct {
	matriculaRepo    domain.MatriculaRepository
	cursoPeriodoRepo domain.CursoPeriodoRepository
}

// Constructor del caso de uso
func NewListarAlumnosMatriculadosUseCase(
	matriculaRepo domain.MatriculaRepository,
	cursoPeriodoRepo domain.CursoPeriodoRepository,
) *ListarAlumnosMatriculadosUseCase {
	return &ListarAlumnosMatriculadosUseCase{
		matriculaRepo:    matriculaRepo,
		cursoPeriodoRepo: cursoPeriodoRepo,
	}
}

func (u *ListarAlumnosMatriculadosUseCase) Execute(periodoID, cursoID int64) shared.APIResponse {

	cursoPeriodo, err := u.cursoPeriodoRepo.FindByPeriodoYCurso(periodoID, cursoID)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el curso en el periodo", nil)
	}
	if !cursoPeriodo.Existe() {
		return shared.NewAPIResponse(404, "El curso en el periodo no existe", nil)
	}

	matriculas, err := u.matriculaRepo.ObtenerMatriculasPorCursoPeriodo(cursoPeriodo.GetID())
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener los alumnos matriculados", nil)
	}

	alumnosDTO := []dto.AlumnoDTO{}
	for _, matricula := range matriculas {
		alumnosDTO = append(alumnosDTO, dto.AlumnoDTO{
			ID:                matricula.GetAlumnoID(),
			Nombre:            matricula.GetAlumno().GetNombre(),
			FechaNacimiento:   matricula.GetAlumno().GetFechaNacimiento(),
			Telefono:          matricula.GetAlumno().GetTelefono(),
			Acudiente:         matricula.GetAlumno().GetAcudiente(),
			AcudienteTelefono: matricula.GetAlumno().GetAcudienteTelefono(),
			Direccion:         matricula.GetAlumno().GetDireccion(),
			Activo:            matricula.GetAlumno().GetActivo(),
			Edad:              matricula.GetAlumno().CalcularEdad(),
		})
	}

	return shared.NewAPIResponse(200, "Alumnos matriculados obtenidos correctamente", alumnosDTO)
}
