package domain

import "ebd/src/view/dto"

type Celebracion struct {
	id         int64
	nombre     string
	repository CelebracionRepository
}

func NewCelebracion(repository CelebracionRepository) *Celebracion {
	return &Celebracion{
		repository: repository,
	}
}

func (c *Celebracion) SetID(id int64) {
	c.id = id
}

func (c *Celebracion) SetNombre(nombre string) {
	c.nombre = nombre
}

func (c *Celebracion) GetID() int64 {
	return c.id
}

func (c *Celebracion) GetNombre() string {
	return c.nombre
}

func (c *Celebracion) Existe() bool {
	return c.id > 0
}

func (c *Celebracion) Crear() error {
	return c.repository.Save(c)
}

func (c *Celebracion) Actualizar() error {
	return c.repository.Update(c)
}

func (c *Celebracion) Eliminar() error {
	return c.repository.Delete(c.id)
}

func (c *Celebracion) ToDTO() dto.CelebracionDto {
	return dto.CelebracionDto{
		ID:     c.id,
		Nombre: c.nombre,
	}
}
