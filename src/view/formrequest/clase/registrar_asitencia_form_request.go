package formrequest

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

type RegistrarAsitenciaFormRequest struct {
	Fecha               string  `json:"fecha" binding:"required"`
	Ofrenda             float64 `json:"ofrenda" binding:"required,gte=0"`
	GrupoID             int64   `json:"grupo_id" binding:"required"`
	ContenidoTematicoID int64   `json:"contenido_tematico_id" binding:"required"`
	AlumnosMatriculados []int64 `json:"alumnos_matriculados" binding:"required,dive,gt=0"`
}

func (r *RegistrarAsitenciaFormRequest) Validate(c *gin.Context) error {
	_, err := time.Parse("2006-01-02", r.Fecha)
	if err != nil {
		return errors.New("la fecha debe tener el formato YYYY-MM-DD")
	}

	if r.Ofrenda < 0 {
		return errors.New("la ofrenda no puede ser un valor negativo")
	}

	if r.GrupoID <= 0 {
		return errors.New("wl ID del grupo debe ser válido")
	}

	if r.ContenidoTematicoID <= 0 {
		return errors.New("wl ID del contenido temático debe ser válido")
	}

	if len(r.AlumnosMatriculados) == 0 {
		return errors.New("debe haber al menos un alumno matriculado")
	}

	return nil
}
