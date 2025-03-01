package domain

import (
	"ebd/src/view/dto"
	"errors"
	"log"
	"time"
)

type Alumno struct {
	id                int64
	nombre            string
	fechaNacimiento   string
	telefono          string
	acudiente         string
	acudienteTelefono string
	direccion         string
	repository        AlumnoRepository
}

func NewAlumno(repository AlumnoRepository) *Alumno {
	return &Alumno{repository: repository}
}

func (a *Alumno) SetID(id int64) {
	a.id = id
}

func (a *Alumno) SetNombre(nombre string) {
	a.nombre = nombre
}

func (a *Alumno) SetFechaNacimiento(fecha string) {
	a.fechaNacimiento = fecha
}

func (a *Alumno) SetTelefono(telefono string) {
	a.telefono = telefono
}

func (a *Alumno) SetAcudiente(acudiente string) {
	a.acudiente = acudiente
}

func (a *Alumno) SetAcudienteTelefono(acudienteTelefono string) {
	a.acudienteTelefono = acudienteTelefono
}

func (a *Alumno) SetDireccion(direccion string) {
	a.direccion = direccion
}

func (a *Alumno) GetID() int64 {
	return a.id
}

func (a *Alumno) GetNombre() string {
	return a.nombre
}

func (a *Alumno) GetFechaNacimiento() string {
	return a.fechaNacimiento
}

func (a *Alumno) GetTelefono() string {
	return a.telefono
}

func (a *Alumno) GetAcudiente() string {
	return a.acudiente
}

func (a *Alumno) GetAcudienteTelefono() string {
	return a.acudienteTelefono
}

func (a *Alumno) GetDireccion() string {
	return a.direccion
}

func (a *Alumno) Existe() bool {
	return a.id > 0
}

func (a *Alumno) Save() error {
	return a.repository.Save(a)
}

func (a *Alumno) Update() error {
	return a.repository.Update(a)
}

func (a *Alumno) Delete() error {
	return a.repository.Delete(a.id)
}

func (a *Alumno) FindByID(id int64) (*Alumno, error) {
	return a.repository.FindByID(id)
}

func (a *Alumno) FindByNombre(nombre string) (*Alumno, error) {
	return a.repository.FindByNombre(nombre)
}

func (a *Alumno) MatricularCurso(cursoPeriodo *CursoPeriodo) error {
	if cursoPeriodo == nil || !cursoPeriodo.Existe() {
		return errors.New("el curso seleccionado no está disponible")
	}

	return a.repository.MatricularCurso(a.id, cursoPeriodo.GetID())
}

func (a *Alumno) TieneCursoMatriculado(periodoID int64) bool {
	return a.repository.TieneCursoMatriculado(a.id, periodoID)
}

func (a *Alumno) CalcularEdad() int {
	fechaNacimiento, err := time.Parse(time.RFC3339, a.fechaNacimiento)
	if err != nil {
		log.Printf("Error al parsear fecha de nacimiento '%s': %v", a.fechaNacimiento, err)
		return 0
	}

	hoy := time.Now()
	edad := hoy.Year() - fechaNacimiento.Year()

	if hoy.YearDay() < fechaNacimiento.YearDay() {
		edad--
	}

	return edad
}

func (a *Alumno) ToDTO() *dto.AlumnoDTO {
	return &dto.AlumnoDTO{
		ID:                a.id,
		Nombre:            a.nombre,
		FechaNacimiento:   a.fechaNacimiento,
		Telefono:          a.telefono,
		Acudiente:         a.acudiente,
		AcudienteTelefono: a.acudienteTelefono,
		Direccion:         a.direccion,
		Edad:              a.CalcularEdad(),
	}
}
