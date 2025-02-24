package formrequest

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

type PeriodoFormRequest struct {
	Nombre      string `json:"nombre" binding:"required"`
	FechaInicio string `json:"fecha_inicio" binding:"required"`
	FechaFin    string `json:"fecha_fin" binding:"required"`
}

func (p *PeriodoFormRequest) Validate(c *gin.Context) error {

	fechaInicio, err := time.Parse("2006-01-02", p.FechaInicio)
	if err != nil {
		return errors.New("fecha_inicio no tiene un formato válido (YYYY-MM-DD)")
	}

	fechaFin, err := time.Parse("2006-01-02", p.FechaFin)
	if err != nil {
		return errors.New("fecha_fin no tiene un formato válido (YYYY-MM-DD)")
	}

	if fechaFin.Before(fechaInicio) {
		return errors.New("fecha_fin debe ser mayor que fecha_inicio")
	}

	return nil
}
