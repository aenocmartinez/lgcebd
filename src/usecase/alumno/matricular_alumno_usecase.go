package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
)

type MatricularAlumnoUseCase struct {
	alumnoRepo       domain.AlumnoRepository
	cursoPeriodoRepo domain.CursoPeriodoRepository
	matriculaRepo    domain.MatriculaRepository
}

func NewMatricularAlumnoUseCase(
	alumnoRepo domain.AlumnoRepository,
	cursoPeriodoRepo domain.CursoPeriodoRepository,
	matriculaRepo domain.MatriculaRepository,
) *MatricularAlumnoUseCase {
	return &MatricularAlumnoUseCase{
		alumnoRepo:       alumnoRepo,
		cursoPeriodoRepo: cursoPeriodoRepo,
		matriculaRepo:    matriculaRepo,
	}
}

func (u *MatricularAlumnoUseCase) Execute(alumnoID, cursoPeriodoID int64) shared.APIResponse {

	alumno, err := u.alumnoRepo.FindByID(alumnoID)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al buscar el alumno", nil)
	}
	if !alumno.Existe() {
		return shared.NewAPIResponse(404, "El alumno no existe", nil)
	}

	cursoPeriodo := u.cursoPeriodoRepo.FindByID(cursoPeriodoID)
	if !cursoPeriodo.Existe() {
		return shared.NewAPIResponse(500, "Curso  no encontrado", nil)
	}
	if !cursoPeriodo.Existe() {
		return shared.NewAPIResponse(404, "El curso en el periodo no existe", nil)
	}

	if u.matriculaRepo.ExisteMatricula(alumnoID, cursoPeriodoID) {
		return shared.NewAPIResponse(400, "El alumno ya está matriculado en este curso", nil)
	}

	if u.alumnoRepo.TieneCursoMatriculado(alumnoID, cursoPeriodo.GetPeriodoID()) {
		return shared.NewAPIResponse(400, "El alumno ya tiene un curso matriculado en este periodo", nil)
	}

	edadAlumno := alumno.CalcularEdad()
	if edadAlumno < cursoPeriodo.GetCurso().GetEdadMinima() || edadAlumno > cursoPeriodo.GetCurso().GetEdadMaxima() {
		return shared.NewAPIResponse(400, "El alumno no cumple con los requisitos de edad para este curso", nil)
	}

	matricula := domain.NewMatriculaEmpty()
	matricula.SetAlumno(alumno)
	matricula.SetCursoPeriodo(cursoPeriodo)

	if err := u.matriculaRepo.Save(matricula); err != nil {
		return shared.NewAPIResponse(500, "Error al matricular al alumno en el curso", nil)
	}

	return shared.NewAPIResponse(200, "Matrícula realizada exitosamente", nil)
}
