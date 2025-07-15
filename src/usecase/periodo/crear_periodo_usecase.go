package usecase

import (
	"ebd/src/domain"
	"ebd/src/shared"
	"ebd/src/view/dto"
	"fmt"
	"sync"

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

	err = u.asociarCursosAlPeriodo(nuevoPeriodo.GetID(), cursos)
	if err != nil {
		return shared.NewAPIResponse(500, "Error al asociar al menos un curso al periodo", nil)
	}

	// Matricular alumnos activos al nuevo periodo de acuerdo a su edad
	alumnos, err := u.alumnoRepo.List()
	if err != nil {
		return shared.NewAPIResponse(500, "Error al obtener los alumnos", nil)
	}

	u.matricularAlumnosAlPeriodo(alumnos)

	return shared.NewAPIResponse(200, "Periodo creado exitosamente.", nuevoPeriodo.ToDTO())
}

func (u *CrearPeriodoUseCase) asociarCursosAlPeriodo(periodoID int64, cursos []dto.CursoDTO) error {
	var wg sync.WaitGroup
	var errOnce sync.Once
	var firstErr error

	for _, curso := range cursos {
		if curso.Estado == "inactivo" {
			continue
		}

		wg.Add(1)
		go func(c dto.CursoDTO) {
			defer wg.Done()
			if err := u.periodoRepo.AgregarCurso(periodoID, c.ID); err != nil {
				errOnce.Do(func() {
					firstErr = err
				})
			}
		}(curso)
	}

	wg.Wait()
	return firstErr
}

func (u *CrearPeriodoUseCase) matricularAlumnosAlPeriodo(alumnos []domain.Alumno) {
	var wg sync.WaitGroup
	matricularUseCase := alumnoUseCase.NewMatricularAlumnoUseCase(u.alumnoRepo, u.cursoPeriodoRepo, u.matriculaRepo)

	for _, alumno := range alumnos {
		if !alumno.GetActivo() {
			continue
		}

		wg.Add(1)
		go func(a domain.Alumno) {
			defer wg.Done()

			edad := a.CalcularEdad()
			periodoID, err := u.cursoPeriodoRepo.ObtenerPeriodoCursoIDPorEdad(edad)
			if err != nil {
				fmt.Printf("Error al obtener el periodo por edad: %v - Alumno ID: %d\n", err, a.GetID())
				u.alumnoRepo.CambiarEstado(a.GetID(), false)
				return
			}

			matricularUseCase.Execute(a.GetID(), periodoID)
		}(alumno)
	}

	wg.Wait()
}
