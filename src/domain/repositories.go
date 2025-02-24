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
}
