package controller

import (
	"ebd/src/shared"
	usecase "ebd/src/usecase/maestro"
	formrequest "ebd/src/view/formrequest/maestro"

	"ebd/src/infraestructure/di"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarMaestros(c *gin.Context) {
	useCase := usecase.NewListarMaestrosUseCase(di.GetContainer().GetMaestroRepository())
	response := useCase.Execute()
	c.JSON(response.StatusCode, response)
}

func CrearMaestro(c *gin.Context) {
	var request formrequest.MaestroFormRequest

	// Validación de datos de entrada
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos de entrada inválidos", nil))
		return
	}
	if err := request.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	useCase := usecase.NewCrearMaestroUseCase(di.GetContainer().GetMaestroRepository())
	response := useCase.Execute(request.ToDTO())
	c.JSON(response.StatusCode, response)
}

func ActualizarMaestro(c *gin.Context) {
	var request formrequest.MaestroFormRequest

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

	useCase := usecase.NewActualizarMaestroUseCase(di.GetContainer().GetMaestroRepository())
	response := useCase.Execute(id, request.ToDTO())
	c.JSON(response.StatusCode, response)
}

func EliminarMaestro(c *gin.Context) {
	id, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
		return
	}

	useCase := usecase.NewEliminarMaestroUseCase(di.GetContainer().GetMaestroRepository())
	response := useCase.Execute(id)
	c.JSON(response.StatusCode, response)
}

func BuscarMaestroPorId(c *gin.Context) {
	id, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
		return
	}

	maestroRepo := di.GetContainer().GetMaestroRepository()
	buscarMaestro := usecase.NewBuscarMaestroUseCase(maestroRepo)
	response := buscarMaestro.Execute(id)
	c.JSON(response.StatusCode, response)
}
