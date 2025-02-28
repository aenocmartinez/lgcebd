package domain

import (
	"ebd/src/view/dto"
)

type CursoPeriodo struct {
	id         int64
	curso      *Curso
	periodo    *Periodo
	repository CursoPeriodoRepository
}

func NewCursoPeriodo(id int64, curso *Curso, periodo *Periodo) *CursoPeriodo {
	return &CursoPeriodo{
		id:      id,
		curso:   curso,
		periodo: periodo,
	}
}

func NewCursoPeriodoEmpty(repository CursoPeriodoRepository) *CursoPeriodo {
	return &CursoPeriodo{
		repository: repository,
	}
}

func (cp *CursoPeriodo) SetRepository(repository CursoPeriodoRepository) {
	cp.repository = repository
}

func (cp *CursoPeriodo) SetID(id int64) {
	cp.id = id
}

func (cp *CursoPeriodo) GetID() int64 {
	return cp.id
}

func (cp *CursoPeriodo) SetCurso(curso *Curso) {
	cp.curso = curso
}

func (cp *CursoPeriodo) GetCurso() *Curso {
	return cp.curso
}

func (cp *CursoPeriodo) SetPeriodo(periodo *Periodo) {
	cp.periodo = periodo
}

func (cp *CursoPeriodo) GetPeriodo() *Periodo {
	return cp.periodo
}

func (cp *CursoPeriodo) GetCursoID() int64 {
	if cp.curso != nil {
		return cp.curso.GetID()
	}
	return 0
}

func (cp *CursoPeriodo) GetPeriodoID() int64 {
	if cp.periodo != nil {
		return cp.periodo.GetID()
	}
	return 0
}

func (cp *CursoPeriodo) Existe() bool {
	return cp.id > 0
}

func (cp *CursoPeriodo) AgregarContenidoTematico(descripcion string) error {
	contenidoTematico := NewContenidoTematico(nil)
	contenidoTematico.SetDescripcion(descripcion)
	return cp.repository.AgregarContenidoTematico(cp.id, contenidoTematico)
}

func (cp *CursoPeriodo) QuitarContenidoTematico(contenidoTematico *ContenidoTematico) error {
	return cp.repository.QuitarContenidoTematico(cp.id, contenidoTematico.GetID())
}

func (cp *CursoPeriodo) ToDTO() dto.ItemCursoPeriodoDTO {
	return dto.ItemCursoPeriodoDTO{
		ID:      cp.id,
		Periodo: *cp.periodo.ToDTO(),
		Curso:   *cp.curso.ToDTO(),
	}
}
