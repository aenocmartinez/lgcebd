package formrequest

import "errors"

type CursoFormRequest struct {
	Nombre     string `json:"nombre" binding:"required"`
	EdadMinima int    `json:"edad_minima" binding:"required"`
	EdadMaxima int    `json:"edad_maxima" binding:"required"`
	Estado     string `json:"estado,omitempty"`
}

func (c *CursoFormRequest) Validate(isUpdate bool) error {
	if c.Nombre == "" {
		return errors.New("el nombre del curso es obligatorio")
	}
	if c.EdadMinima <= 0 {
		return errors.New("la edad mínima debe ser mayor a 0")
	}
	if c.EdadMaxima <= 0 {
		return errors.New("la edad máxima debe ser mayor a 0")
	}
	if c.EdadMinima >= c.EdadMaxima {
		return errors.New("la edad mínima debe ser menor que la edad máxima")
	}
	if isUpdate && c.Estado == "" {
		return errors.New("el estado es obligatorio en la actualización")
	}
	return nil
}
