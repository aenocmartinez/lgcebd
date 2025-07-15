package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
	"fmt"

	alumnoUseCase "ebd/src/usecase/alumno"
)

type CrearPeriodoUseCase struct {
	periodoRepo      domain.PeriodoRepository
	cursoRepo        domain.CursoRepository
	matriculaRepo    domain.MatriculaRepository
	alumnoRepo       domain.AlumnoRepository
	cursoPeriodoRepo domain.CursoPeriodoRepository
}

func NewCrearPeriodoUseCase(
	periodoRepo domain.PeriodoRepository,
	cursoRepo domain.CursoRepository,
	matriculaRepo domain.MatriculaRepository,
	alumnoRepo domain.AlumnoRepository,
	cursoPeriodoRepo domain.CursoPeriodoRepository,
) *CrearPeriodoUseCase {
	return &CrearPeriodoUseCase{
		periodoRepo:      periodoRepo,
		cursoRepo:        cursoRepo,
		matriculaRepo:    matriculaRepo,
		alumnoRepo:       alumnoRepo,
		cursoPeriodoRepo: cursoPeriodoRepo,
	}
}

func (u *CrearPeriodoUseCase) Execute(request dto.PeriodoDTO) shared.APIResponse {

	periodo, err := u.periodoRepo.FindByNombre(request.Nombre)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al verificar la existencia del periodo", nil)
	}

	if periodo.Existe() {
		return shared.NewAPIResponse(400, "Ya existe un periodo con ese nombre", nil)
	}

	nuevoPeriodo := domain.NewPeriodo(u.periodoRepo)
	nuevoPeriodo.SetNombre(request.Nombre)
	nuevoPeriodo.SetFechaInicio(request.FechaInicio)
	nuevoPeriodo.SetFechaFin(request.FechaFin)

	err = u.periodoRepo.Save(nuevoPeriodo)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al crear el periodo", nil)
	}

	// Asociar cursos al nuevo periodo
	cursos, err := u.cursoRepo.List()
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener los cursos", nil)
	}

	if len(cursos) == 0 {
		return shared.NewAPIResponse(200, "Periodo creado sin cursos asociados, no hay cursos habilitados.", nil)
	}

	for _, curso := range cursos {

		if curso.Estado == "inactivo" {
			continue
		}

		err := u.periodoRepo.AgregarCurso(nuevoPeriodo.GetID(), curso.ID)
		if err != nil {
			return shared.NewAPIResponse(500, "Error al asociar el curso al periodo", nil)
		}
	}

	// Matricular alumnos activos al nuevo periodo de acuerdo a su edad
	alumnos, err := u.alumnoRepo.List()
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener los alumnos", nil)
	}

	matricularUseCase := alumnoUseCase.NewMatricularAlumnoUseCase(u.alumnoRepo, u.cursoPeriodoRepo, u.matriculaRepo)

	for _, alumno := range alumnos {

		if !alumno.GetActivo() {
			continue
		}

		periodoID, err := u.cursoPeriodoRepo.ObtenerPeriodoCursoIDPorEdad(alumno.CalcularEdad())
		if err != nil {
			fmt.Println("Error al obtener el periodo por edad:", err.Error(), " - Alumno ID:", alumno.GetID())
			continue
		}

		matricularUseCase.Execute(alumno.GetID(), periodoID)
	}

	return shared.NewAPIResponse(200, "Periodo creado y cursos asociados exitosamente.", nuevoPeriodo.ToDTO())
}
