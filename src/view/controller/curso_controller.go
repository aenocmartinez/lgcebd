package controller

import (
	"net/http"
	"strconv"

	"ebd/src/shared"
	usecase "ebd/src/usecase/curso"
	formrequest "ebd/src/view/formrequest/curso"

	"github.com/gin-gonic/gin"
)

func ListarCursos(c *gin.Context) {
	listarCursos := usecase.ListarCursosUseCase{}
	response := listarCursos.Execute()
	c.JSON(response.StatusCode, response)
}

func CrearCurso(c *gin.Context) {
	var request formrequest.CursoFormRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos inválidos", nil))
		return
	}

	crearCurso := usecase.CrearCursoUseCase{}
	response := crearCurso.Execute(request)
	c.JSON(response.StatusCode, response)
}

func ActualizarCurso(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
		return
	}

	var request formrequest.CursoFormRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos inválidos", nil))
		return
	}

	actualizarCurso := usecase.UpdateCursoUseCase{}
	response := actualizarCurso.Execute(id, request)
	c.JSON(response.StatusCode, response)
}

func EliminarCurso(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
		return
	}

	eliminarCurso := usecase.EliminarCursoUseCase{}
	response := eliminarCurso.Execute(id)
	c.JSON(response.StatusCode, response)
}
