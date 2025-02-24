package domain

import "ebd/src/view/dto"

type Periodo struct {
	id          int64
	nombre      string
	fechaInicio string
	fechaFin    string
	repository  PeriodoRepository
}

func NewPeriodo(repository PeriodoRepository) *Periodo {
	return &Periodo{repository: repository}
}

func (p *Periodo) SetID(id int64) {
	p.id = id
}

func (p *Periodo) SetNombre(nombre string) {
	p.nombre = nombre
}

func (p *Periodo) SetFechaInicio(fechaInicio string) {
	p.fechaInicio = fechaInicio
}

func (p *Periodo) SetFechaFin(fechaFin string) {
	p.fechaFin = fechaFin
}

func (p *Periodo) GetID() int64 {
	return p.id
}

func (p *Periodo) GetNombre() string {
	return p.nombre
}

func (p *Periodo) GetFechaInicio() string {
	return p.fechaInicio
}

func (p *Periodo) GetFechaFin() string {
	return p.fechaFin
}

func (p *Periodo) Save() error {
	return p.repository.Save(p)
}

func (p *Periodo) Update() error {
	return p.repository.Update(p)
}

func (p *Periodo) Delete() error {
	return p.repository.Delete(p.id)
}

func (p *Periodo) FindByID(id int64) (*Periodo, error) {
	return p.repository.FindByID(id)
}

func (p *Periodo) FindByNombre(nombre string) (*Periodo, error) {
	return p.repository.FindByNombre(nombre)
}

func (p *Periodo) Existe() bool {
	return p.id > 0
}

// Convertir a DTO
func (p *Periodo) ToDTO() *dto.PeriodoDTO {
	return &dto.PeriodoDTO{
		ID:          p.id,
		Nombre:      p.nombre,
		FechaInicio: p.fechaInicio,
		FechaFin:    p.fechaFin,
	}
}
