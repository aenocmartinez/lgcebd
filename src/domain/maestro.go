package domain

import "ebd/src/view/dto"

type Maestro struct {
	id              int64
	nombre          string
	telefono        string
	fechaNacimiento string
	estado          string
	repository      MaestroRepository
}

func NewMaestro(repository MaestroRepository) *Maestro {
	return &Maestro{repository: repository}
}

func (m *Maestro) SetID(id int64) {
	m.id = id
}

func (m *Maestro) GetID() int64 {
	return m.id
}

func (m *Maestro) SetNombre(nombre string) {
	m.nombre = nombre
}

func (m *Maestro) GetNombre() string {
	return m.nombre
}

func (m *Maestro) SetTelefono(telefono string) {
	m.telefono = telefono
}

func (m *Maestro) GetTelefono() string {
	return m.telefono
}

func (m *Maestro) SetFechaNacimiento(fecha string) {
	m.fechaNacimiento = fecha
}

func (m *Maestro) GetFechaNacimiento() string {
	return m.fechaNacimiento
}

func (m *Maestro) SetEstado(estado string) {
	m.estado = estado
}

func (m *Maestro) GetEstado() string {
	return m.estado
}

func (m *Maestro) Activar() {
	m.estado = "activo"
}

func (m *Maestro) Desactivar() {
	m.estado = "inactivo"
}

func (m *Maestro) Existe() bool {
	return m.id > 0
}

func (m *Maestro) ToDTO() *dto.MaestroDTO {
	return &dto.MaestroDTO{
		ID:              m.id,
		Nombre:          m.nombre,
		Telefono:        m.telefono,
		FechaNacimiento: m.fechaNacimiento,
		Estado:          m.estado,
	}
}
