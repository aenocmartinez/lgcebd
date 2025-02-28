package domain

import "ebd/src/view/dto"

type UserRepository interface {
	FindByID(id int64) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	Save(user *User) error
	Update(user *User) error
	Delete(id int64) error
}

type CursoRepository interface {
	FindByID(id int64) (*Curso, error)
	FindByNombre(nombre string) (*Curso, error)
	Save(curso *Curso) error
	Update(curso *Curso) error
	Delete(id int64) error
	List() ([]dto.CursoDTO, error)
}

type PeriodoRepository interface {
	FindByID(id int64) (*Periodo, error)
	FindByNombre(nombre string) (*Periodo, error)
	List() ([]dto.PeriodoDTO, error)
	Save(periodo *Periodo) error
	Update(periodo *Periodo) error
	Delete(id int64) error
	AgregarCurso(periodoID, cursoID int64) error
	ObtenerCursos(periodoID int64) ([]dto.CursoPeriodoDTO, error)
}

type AlumnoRepository interface {
	FindByID(id int64) (*Alumno, error)
	FindByNombre(nombre string) (*Alumno, error)
	List() ([]dto.AlumnoDTO, error)
	Save(alumno *Alumno) error
	Update(alumno *Alumno) error
	Delete(id int64) error
	MatricularCurso(alumnoID, cursoPeriodoID int64) error
	TieneCursoMatriculado(alumnoID, periodoID int64) bool
}

type CursoPeriodoRepository interface {
	FindByPeriodoYCurso(periodoID, cursoID int64) (*CursoPeriodo, error)
	ObtenerCursosPorPeriodo(periodoID int64) ([]dto.CursoPeriodoDTO, error)
	FindByID(id int64) *CursoPeriodo
}

type MatriculaRepository interface {
	Save(matricula *Matricula) error
	ExisteMatricula(alumnoID, cursoPeriodoID int64) bool
	ObtenerMatriculasPorAlumno(alumnoID int64) ([]Matricula, error)
	ObtenerMatriculasPorCursoPeriodo(cursoPeriodoID int64) ([]Matricula, error)
	Delete(id int64) error
}

type MaestroRepository interface {
	FindByID(id int64) *Maestro
	List() ([]Maestro, error)
	Save(maestro *Maestro) error
	Update(maestro *Maestro) error
	Delete(id int64) error
}

type CelebracionRepository interface {
	FindByID(id int64) *Celebracion
	FindByNombre(nombre string) *Celebracion
	Save(celebracion *Celebracion) error
	Update(celebracion *Celebracion) error
	Delete(id int64) error
	List() []Celebracion
}

type GrupoRepository interface {
	FindByID(id int64) *Grupo
	FindByCursoPeriodoYCelebracion(cursoPeriodoID, celebracionID int64) *Grupo
	Save(grupo *Grupo) error
	Update(grupo *Grupo) error
	Delete(id int64) error
	List() []Grupo
	AgregarMaestro(grupoID int64, maestroID int64) error
	QuitarMaestros(grupoID int64) error
	ListarMaestros(grupoID int64) []GrupoMaestro
}
