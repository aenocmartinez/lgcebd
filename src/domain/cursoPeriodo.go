package domain

type CursoPeriodo struct {
	id      int64
	curso   *Curso
	periodo *Periodo
}

func NewCursoPeriodo(id int64, curso *Curso, periodo *Periodo) *CursoPeriodo {
	return &CursoPeriodo{
		id:      id,
		curso:   curso,
		periodo: periodo,
	}
}

func NewCursoPeriodoEmpty() *CursoPeriodo {
	return &CursoPeriodo{}
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
