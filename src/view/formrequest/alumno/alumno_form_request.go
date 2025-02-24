package formrequest

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

type AlumnoFormRequest struct {
	Nombre            string `json:"nombre" binding:"required"`
	FechaNacimiento   string `json:"fecha_nacimiento" binding:"required"`
	Telefono          string `json:"telefono" binding:"required,len=10"`
	Acudiente         string `json:"acudiente" binding:"required"`
	AcudienteTelefono string `json:"acudiente_telefono" binding:"required,len=10"`
	Direccion         string `json:"direccion" binding:"required"`
}

func (a *AlumnoFormRequest) Validate(c *gin.Context) error {
	_, err := time.Parse("2006-01-02", a.FechaNacimiento)
	if err != nil {
		return errors.New("la fecha de nacimiento debe tener el formato YYYY-MM-DD")
	}

	fechaNacimiento, _ := time.Parse("2006-01-02", a.FechaNacimiento)
	if fechaNacimiento.After(time.Now()) {
		return errors.New("la fecha de nacimiento no puede estar en el futuro")
	}

	return nil
}
