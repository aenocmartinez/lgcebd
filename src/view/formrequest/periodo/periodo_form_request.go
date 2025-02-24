package formrequest

import (
	"ebd/src/view/dto"
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
	if p.Nombre == "" || p.FechaInicio == "" || p.FechaFin == "" {
		return errors.New("todos los campos son obligatorios")
	}

	fechaInicio, err := time.Parse("2006-01-02", p.FechaInicio)
	if err != nil {
		return errors.New("fecha de inicio inválida, formato esperado: YYYY-MM-DD")
	}

	fechaFin, err := time.Parse("2006-01-02", p.FechaFin)
	if err != nil {
		return errors.New("fecha de fin inválida, formato esperado: YYYY-MM-DD")
	}

	if fechaFin.Before(fechaInicio) {
		return errors.New("la fecha de fin debe ser posterior a la fecha de inicio")
	}

	return nil
}

func (p *PeriodoFormRequest) ToDTO() dto.PeriodoDTO {
	return dto.PeriodoDTO{
		Nombre:      p.Nombre,
		FechaInicio: p.FechaInicio,
		FechaFin:    p.FechaFin,
	}
}
