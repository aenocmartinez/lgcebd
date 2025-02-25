package formrequest

import (
	"ebd/src/view/dto"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type MaestroFormRequest struct {
	Nombre          string `json:"nombre" binding:"required"`
	Telefono        string `json:"telefono" binding:"required"`
	FechaNacimiento string `json:"fecha_nacimiento" binding:"required"`
	Estado          string `json:"estado" binding:"omitempty,oneof=activo inactivo"`
}

func (r *MaestroFormRequest) Validate(c *gin.Context) error {
	r.Nombre = strings.TrimSpace(r.Nombre)
	if r.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}

	r.Telefono = strings.TrimSpace(r.Telefono)
	if r.Telefono == "" {
		return errors.New("el teléfono es obligatorio")
	}

	_, err := time.Parse("2006-01-02", r.FechaNacimiento)
	if err != nil {
		return errors.New("la fecha de nacimiento debe tener el formato YYYY-MM-DD")
	}

	if r.Estado != "" && r.Estado != "activo" && r.Estado != "inactivo" {
		return errors.New("el estado solo puede ser 'activo' o 'inactivo'")
	}

	return nil
}

func (r *MaestroFormRequest) ToDTO() dto.MaestroDTO {
	return dto.MaestroDTO{
		Nombre:          r.Nombre,
		Telefono:        r.Telefono,
		FechaNacimiento: r.FechaNacimiento,
		Estado:          r.Estado,
	}
}
