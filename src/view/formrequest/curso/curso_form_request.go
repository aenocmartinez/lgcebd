package formrequest

import (
	"errors"
)

type CursoFormRequest struct {
	Nombre     string `json:"nombre"`
	EdadMinima int    `json:"edad_minima"`
	EdadMaxima int    `json:"edad_maxima"`
}

func (r *CursoFormRequest) Validate() error {
	if r.Nombre == "" {
		return errors.New("el nombre del curso es obligatorio")
	}
	if r.EdadMinima <= 0 {
		return errors.New("la edad mínima debe ser mayor a 0")
	}
	if r.EdadMaxima <= 0 {
		return errors.New("la edad máxima debe ser mayor a 0")
	}
	if r.EdadMaxima <= r.EdadMinima {
		return errors.New("la edad máxima debe ser mayor que la edad mínima")
	}
	return nil
}
