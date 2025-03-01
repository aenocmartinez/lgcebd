package domain

type Asistencia struct {
	id         int64
	clase      *Clase
	matricula  *Matricula
	repository ClaseRepository
}

func NewAsistencia(repository ClaseRepository) *Asistencia {
	return &Asistencia{
		repository: repository,
	}
}

func (a *Asistencia) SetID(id int64) {
	a.id = id
}

func (a *Asistencia) SetClase(clase *Clase) {
	a.clase = clase
}

func (a *Asistencia) SetMatricula(matricula *Matricula) {
	a.matricula = matricula
}

func (a *Asistencia) GetID() int64 {
	return a.id
}

func (a *Asistencia) GetClase() *Clase {
	return a.clase
}

func (a *Asistencia) GetMatricula() *Matricula {
	return a.matricula
}
