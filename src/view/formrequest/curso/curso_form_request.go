package formrequest

import (
	"ebd/src/view/dto"
	"errors"

	"github.com/gin-gonic/gin"
)

type CursoFormRequest struct {
	Nombre     string `json:"nombre" binding:"required"`
	EdadMinima int    `json:"edad_minima" binding:"required"`
	EdadMaxima int    `json:"edad_maxima" binding:"required"`
	Estado     string `json:"estado,omitempty"` // No es obligatorio en creación
}

func (c *CursoFormRequest) Validate(ctx *gin.Context) error {
	if c.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}
	if c.EdadMinima < 0 {
		return errors.New("la edad mínima no puede ser negativa")
	}
	if c.EdadMaxima < c.EdadMinima {
		return errors.New("la edad máxima debe ser mayor o igual a la edad mínima")
	}
	return nil
}

func (c *CursoFormRequest) ToDTO() dto.CursoDTO {
	return dto.CursoDTO{
		Nombre:     c.Nombre,
		EdadMinima: c.EdadMinima,
		EdadMaxima: c.EdadMaxima,
		Estado:     c.Estado,
	}
}
