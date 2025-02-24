package controller

import (
	"ebd/src/shared"
	usecase "ebd/src/usecase/periodo"
	formrequest "ebd/src/view/formrequest/periodo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarPeriodos(c *gin.Context) {
	listarPeriodos := usecase.ListarPeriodosUseCase{}
	response := listarPeriodos.Execute()
	c.JSON(response.StatusCode, response)
}

func CrearPeriodo(c *gin.Context) {
	var request formrequest.PeriodoFormRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos de entrada inválidos", nil))
		return
	}

	if err := request.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	crearPeriodo := usecase.CrearPeriodoUseCase{}
	response := crearPeriodo.Execute(request)
	c.JSON(response.StatusCode, response)
}

func ActualizarPeriodo(c *gin.Context) {
	var request formrequest.PeriodoFormRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos de entrada inválidos", nil))
		return
	}

	if err := request.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	id, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
		return
	}

	actualizarPeriodo := usecase.ActualizarPeriodoUseCase{}
	response := actualizarPeriodo.Execute(id, request)
	c.JSON(response.StatusCode, response)
}

func EliminarPeriodo(c *gin.Context) {

	id, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
		return
	}

	eliminarPeriodo := usecase.EliminarPeriodoUseCase{}
	response := eliminarPeriodo.Execute(id)
	c.JSON(response.StatusCode, response)
}
