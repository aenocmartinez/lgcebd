package controller

import (
	"ebd/src/infraestructure/di"
	"ebd/src/shared"
	usecase "ebd/src/usecase/periodo"
	formrequest "ebd/src/view/formrequest/periodo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarPeriodos(c *gin.Context) {
	useCase := usecase.NewListarPeriodosUseCase(di.GetContainer().GetPeriodoRepository())
	response := useCase.Execute()
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

	useCase := usecase.NewCrearPeriodoUseCase(di.GetContainer().GetPeriodoRepository(), di.GetContainer().GetCursoRepository())
	response := useCase.Execute(request.ToDTO())
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

	useCase := usecase.NewActualizarPeriodoUseCase(di.GetContainer().GetPeriodoRepository())
	response := useCase.Execute(id, request.ToDTO())
	c.JSON(response.StatusCode, response)
}

func EliminarPeriodo(c *gin.Context) {
	id, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
		return
	}

	useCase := usecase.NewEliminarPeriodoUseCase(di.GetContainer().GetPeriodoRepository())
	response := useCase.Execute(id)
	c.JSON(response.StatusCode, response)
}
