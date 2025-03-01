package domain

import "ebd/src/view/dto"

type Clase struct {
	id                int64
	fecha             string
	ofrenda           float64
	grupo             *Grupo
	contenidoTematico *ContenidoTematico
	repository        ClaseRepository
}

func NewClase(repository ClaseRepository) *Clase {
	return &Clase{
		repository: repository,
	}
}

func (c *Clase) SetID(id int64) {
	c.id = id
}

func (c *Clase) GetID() int64 {
	return c.id
}

func (c *Clase) SetFecha(fecha string) {
	c.fecha = fecha
}

func (c *Clase) GetFecha() string {
	return c.fecha
}

func (c *Clase) SetOfrenda(ofrenda float64) {
	c.ofrenda = ofrenda
}

func (c *Clase) GetOfrenda() float64 {
	return c.ofrenda
}

func (c *Clase) SetGrupo(grupo *Grupo) {
	c.grupo = grupo
}

func (c *Clase) GetGrupo() *Grupo {
	return c.grupo
}

func (c *Clase) SetContenidoTematico(contenidoTematico *ContenidoTematico) {
	c.contenidoTematico = contenidoTematico
}

func (c *Clase) GetContenidoTematico() *ContenidoTematico {
	return c.contenidoTematico
}

func (c *Clase) Existe() bool {
	return c.id > 0
}

func (c *Clase) Crear() error {
	return c.repository.Save(c)
}

func (c *Clase) Actualizar() error {
	return c.repository.Update(c)
}

func (c *Clase) Eliminar() error {
	return c.repository.Delete(c.id)
}

func (c *Clase) RegistrarAsistencia(matricula *Matricula) error {
	return c.repository.RegistrarAsistencia(c.id, matricula.id)
}

func (c *Clase) ToDTO() dto.ClaseDTO {
	return dto.ClaseDTO{
		ID:                c.id,
		Fecha:             c.fecha,
		Ofreda:            c.ofrenda,
		Grupo:             c.grupo.ToDTO(),
		ContenidoTematico: c.contenidoTematico.ToDTO(),
	}
}
