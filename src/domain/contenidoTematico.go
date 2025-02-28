package domain

import "ebd/src/view/dto"

type ContenidoTematico struct {
	id           int64
	descripcion  string
	cursoPeriodo *CursoPeriodo
	repository   ContenidoTematicoRepository
}

func NewContenidoTematico(repository ContenidoTematicoRepository) *ContenidoTematico {
	return &ContenidoTematico{
		repository: repository,
	}
}

func NewContenidoTematicoEmpty() *ContenidoTematico {
	return &ContenidoTematico{}
}

func (c *ContenidoTematico) SetID(id int64) {
	c.id = id
}

func (c *ContenidoTematico) GetID() int64 {
	return c.id
}

func (c *ContenidoTematico) SetDescripcion(descripcion string) {
	c.descripcion = descripcion
}

func (c *ContenidoTematico) GetDescripcion() string {
	return c.descripcion
}

func (c *ContenidoTematico) SetCursoPeriodo(cursoPeriodo *CursoPeriodo) {
	c.cursoPeriodo = cursoPeriodo
}

func (c *ContenidoTematico) GetCursoPeriodo() *CursoPeriodo {
	return c.cursoPeriodo
}

func (c *ContenidoTematico) Existe() bool {
	return c.id > 0
}

func (c *ContenidoTematico) Crear() error {
	return c.repository.Save(c)
}

func (c *ContenidoTematico) Actualizar() error {
	return c.repository.Update(c)
}

func (c *ContenidoTematico) Eliminar() error {
	return c.repository.Delete(c.id)
}

func (c *ContenidoTematico) ToDTO() dto.ConteniodoTematicoDTO {
	return dto.ConteniodoTematicoDTO{
		ID:           c.id,
		Descripcion:  c.descripcion,
		CursoPeriodo: c.cursoPeriodo.ToDTO(),
	}
}
