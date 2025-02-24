package domain

import "ebd/src/view/dto"

type Curso struct {
	id         int64
	nombre     string
	edadMinima int
	edadMaxima int
	estado     string
	repository CursoRepository
}

func NewCurso(repository CursoRepository) *Curso {
	return &Curso{repository: repository}
}

func (c *Curso) SetID(id int64) {
	c.id = id
}

func (c *Curso) SetNombre(nombre string) {
	c.nombre = nombre
}

func (c *Curso) SetEdadMinima(edadMinima int) {
	c.edadMinima = edadMinima
}

func (c *Curso) SetEdadMaxima(edadMaxima int) {
	c.edadMaxima = edadMaxima
}

func (c *Curso) SetEstado(estado string) {
	c.estado = estado
}

func (c *Curso) GetID() int64 {
	return c.id
}

func (c *Curso) GetNombre() string {
	return c.nombre
}

func (c *Curso) GetEdadMinima() int {
	return c.edadMinima
}

func (c *Curso) GetEdadMaxima() int {
	return c.edadMaxima
}

func (c *Curso) GetEstado() string {
	return c.estado
}

// Métodos de persistencia
func (c *Curso) Save() error {
	return c.repository.Save(c)
}

func (c *Curso) Update() error {
	return c.repository.Update(c)
}

func (c *Curso) Delete() error {
	return c.repository.Delete(c.id)
}

func (c *Curso) FindByID(id int64) (*Curso, error) {
	return c.repository.FindByID(id)
}

func (c *Curso) Existe() bool {
	return c.id > 0
}

func (c *Curso) ToDTO() *dto.CursoDTO {
	return &dto.CursoDTO{
		ID:         c.id,
		Nombre:     c.nombre,
		EdadMinima: c.edadMinima,
		EdadMaxima: c.edadMaxima,
		Estado:     c.estado,
	}
}
