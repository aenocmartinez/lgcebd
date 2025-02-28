package formrequest

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type CelebracionFormRequest struct {
	Nombre string `json:"nombre" binding:"required,max=255"`
}

func (c *CelebracionFormRequest) Validate(ctx *gin.Context) error {
	if len(c.Nombre) == 0 {
		return errors.New("el nombre es obligatorio")
	}
	if len(c.Nombre) > 255 {
		return errors.New("el nombre no puede exceder los 255 caracteres")
	}
	return nil
}
