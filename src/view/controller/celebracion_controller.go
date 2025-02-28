package controller

import (
	"ebd/src/infraestructure/di"
	"ebd/src/shared"
	usecase "ebd/src/usecase/celebracion"
	formrequest "ebd/src/view/formrequest/celebracion"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarCelebraciones(c *gin.Context) {

	listarCelebraciones := usecase.NewListarCelebracionUseCase(di.GetContainer().GetCelebracionRepository())
	response := listarCelebraciones.Execute()

	c.JSON(response.StatusCode, response)
}

func CrearCelebracion(c *gin.Context) {

	var request formrequest.CelebracionFormRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos de entrada inválidos", err.Error()))
		return
	}

	if err := request.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	crearCelebracion := usecase.NewCrearCelabracionUseCase(di.GetContainer().GetCelebracionRepository())
	response := crearCelebracion.Execute(request.Nombre)
	c.JSON(response.StatusCode, response)
}

func ActualizarCelebracion(c *gin.Context) {

	celebracionID, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID Inválido", nil))
		return
	}

	var request formrequest.CelebracionFormRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos de entrada inválidos", nil))
		return
	}

	if err := request.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	actualizarCelebracion := usecase.NewActualizarCelebracionUseCase(di.GetContainer().GetCelebracionRepository())
	response := actualizarCelebracion.Execute(celebracionID, request.Nombre)
	c.JSON(response.StatusCode, response)
}

func EliminarCelebracion(c *gin.Context) {
	celebracionID, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID Inválido", nil))
		return
	}

	eliminarCelebracion := usecase.NewEliminarCelebracionUseCase(di.GetContainer().GetCelebracionRepository())
	response := eliminarCelebracion.Execute(celebracionID)
	c.JSON(response.StatusCode, response)
}
