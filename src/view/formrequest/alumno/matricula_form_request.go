package formrequest

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type MatriculaFormRequest struct {
	CursoPeriodoID int64 `json:"curso_periodo_id" binding:"required"`
}

func (r *MatriculaFormRequest) Validate(c *gin.Context) error {

	if r.CursoPeriodoID <= 0 {
		return errors.New("el ID del curso en el periodo es obligatorio y debe ser válido")
	}

	return nil
}
