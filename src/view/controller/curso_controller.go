package controller

import (
	"ebd/src/usecase/curso"
	formrequest "ebd/src/view/formrequest/curso"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarCursos(c *gin.Context) {

	listarCursosUseCase := curso.ListarCursosUseCase{}

	response := listarCursosUseCase.Execute()

	c.JSON(response.StatusCode, response)
}

func CrearCurso(c *gin.Context) {
	var request formrequest.CursoFormRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos"})
		return
	}

	crearCursoUseCase := curso.CrearCursoUseCase{}
	response := crearCursoUseCase.Execute(request)

	c.JSON(response.StatusCode, response)
}
