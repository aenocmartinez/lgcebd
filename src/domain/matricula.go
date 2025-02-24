package domain

type Matricula struct {
	id           int64
	alumno       *Alumno
	cursoPeriodo *CursoPeriodo
}

func NewMatricula(id int64, alumno *Alumno, cursoPeriodo *CursoPeriodo) *Matricula {
	return &Matricula{
		id:           id,
		alumno:       alumno,
		cursoPeriodo: cursoPeriodo,
	}
}

func NewMatriculaEmpty() *Matricula {
	return &Matricula{}
}

func (m *Matricula) SetID(id int64) {
	m.id = id
}

func (m *Matricula) GetID() int64 {
	return m.id
}

func (m *Matricula) SetAlumno(alumno *Alumno) {
	m.alumno = alumno
}

func (m *Matricula) GetAlumno() *Alumno {
	return m.alumno
}

func (m *Matricula) SetCursoPeriodo(cursoPeriodo *CursoPeriodo) {
	m.cursoPeriodo = cursoPeriodo
}

func (m *Matricula) GetCursoPeriodo() *CursoPeriodo {
	return m.cursoPeriodo
}

func (m *Matricula) GetAlumnoID() int64 {
	if m.alumno != nil {
		return m.alumno.GetID()
	}
	return 0
}

func (m *Matricula) GetCursoPeriodoID() int64 {
	if m.cursoPeriodo != nil {
		return m.cursoPeriodo.GetID()
	}
	return 0
}

func (m *Matricula) Existe() bool {
	return m.id > 0
}
