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
	List() ([]Alumno, error)
	Save(alumno *Alumno) error
	Update(alumno *Alumno) error
	Delete(id int64) error
	MatricularCurso(alumnoID, cursoPeriodoID int64) error
	TieneCursoMatriculado(alumnoID, periodoID int64) bool
	CambiarEstado(alumnoID int64, nuevoEstado bool) error
}

type CursoPeriodoRepository interface {
	FindByPeriodoYCurso(periodoID, cursoID int64) (*CursoPeriodo, error)
	ObtenerCursosPorPeriodo(periodoID int64) ([]dto.CursoPeriodoDTO, error)
	AgregarContenidoTematico(cursoPeriodoID int64, contenidoTematico *ContenidoTematico) error
	QuitarContenidoTematico(cursoPeriodoID int64, contenidoTematicoID int64) error
	ListarContenidoTematico(cursoPeriodoID int64) []ContenidoTematico
	FindByID(id int64) *CursoPeriodo
	ObtenerPeriodoCursoIDPorEdad(edad int) (int64, error)
}

type MatriculaRepository interface {
	Save(matricula *Matricula) error
	ExisteMatricula(alumnoID, cursoPeriodoID int64) bool
	ObtenerMatriculasPorAlumno(alumnoID int64) ([]Matricula, error)
	ObtenerMatriculasPorCursoPeriodo(cursoPeriodoID int64) ([]Matricula, error)
	Delete(id int64) error
	FindByID(matriculaID int64) (*Matricula, error)
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

type ContenidoTematicoRepository interface {
	FindByID(id int64) *ContenidoTematico
	FindByDescripcion(cursoPeriodoID int64, descripcion string) *ContenidoTematico
	Save(conteniodoTematico *ContenidoTematico) error
	Update(conteniodoTematico *ContenidoTematico) error
	Delete(id int64) error
	List() []ContenidoTematico
}

type ClaseRepository interface {
	FindByID(id int64) *Clase
	FindByGrupoFecha(grupoID int64, fecha string) *Clase
	List() ([]Clase, error)
	Save(clase *Clase) error
	Update(clase *Clase) error
	Delete(claseID int64) error
	RegistrarAsistencia(claseID int64, matriculaID int64) error
}
