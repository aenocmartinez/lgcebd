package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type CambiarEstadoAlumnoUseCase struct {
	alumnoRepo       domain.AlumnoRepository
	cursoPeriodoRepo domain.CursoPeriodoRepository
	matriculaRepo    domain.MatriculaRepository
}

func NewCambiarEstadoAlumnoUseCase(
	alumnoRepo domain.AlumnoRepository,
	cursoPeriodoRepo domain.CursoPeriodoRepository,
	matriculaRepo domain.MatriculaRepository,
) *CambiarEstadoAlumnoUseCase {
	return &CambiarEstadoAlumnoUseCase{
		alumnoRepo:       alumnoRepo,
		cursoPeriodoRepo: cursoPeriodoRepo,
		matriculaRepo:    matriculaRepo,
	}
}

func (uc *CambiarEstadoAlumnoUseCase) CambiarEstado(alumnoID int64) shared.APIResponse {
	alumno, err := uc.alumnoRepo.FindByID(alumnoID)
	if err != nil {
		return shared.NewAPIResponse(404, "Alumno no encontrado", err.Error())
	}

	if !alumno.Existe() {
		return shared.NewAPIResponse(404, "Alumno no encontrado", nil)
	}

	var nuevoEstado bool = !alumno.GetActivo()

	err = uc.alumnoRepo.CambiarEstado(alumnoID, nuevoEstado)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al cambiar el estado del alumno", nil)
	}

	cursoPeriodoID, err := uc.cursoPeriodoRepo.ObtenerPeriodoCursoIDPorEdad(alumno.CalcularEdad())
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener el curso periodo por edad", err.Error())
	}

	if nuevoEstado {

		existe := uc.matriculaRepo.ExisteMatricula(alumnoID, cursoPeriodoID)
		if existe {
			return shared.NewAPIResponse(200, "Estado del alumno cambiado exitosamente", nil)
		}

		matricularUseCase := NewMatricularAlumnoUseCase(uc.alumnoRepo, uc.cursoPeriodoRepo, uc.matriculaRepo)

		response := matricularUseCase.Execute(alumnoID, cursoPeriodoID)

		if response.StatusCode != 200 {
			return shared.NewAPIResponse(response.StatusCode, "Error al matricular al alumno en el curso", response.Data)
		}

	}

	return shared.NewAPIResponse(200, "Estado del alumno cambiado exitosamente", nil)
}
