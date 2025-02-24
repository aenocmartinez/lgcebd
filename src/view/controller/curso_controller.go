package controller

import (
	"ebd/src/infraestructure/di"
	"ebd/src/shared"
	usecase "ebd/src/usecase/curso"
	formrequest "ebd/src/view/formrequest/curso"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarCursos(c *gin.Context) {
	useCase := usecase.NewListarCursosUseCase(di.GetContainer().GetCursoRepository())
	response := useCase.Execute()
	c.JSON(response.StatusCode, response)
}

func CrearCurso(c *gin.Context) {
	var request formrequest.CursoFormRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos de entrada inválidos", nil))
		return
	}

	if err := request.Validate(c); err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	useCase := usecase.NewCrearCursoUseCase(di.GetContainer().GetCursoRepository())
	response := useCase.Execute(request.ToDTO())
	c.JSON(response.StatusCode, response)
}

func ActualizarCurso(c *gin.Context) {
	var request formrequest.CursoFormRequest

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

	useCase := usecase.NewActualizarCursoUseCase(di.GetContainer().GetCursoRepository())
	response := useCase.Execute(id, request.ToDTO())
	c.JSON(response.StatusCode, response)
}

func EliminarCurso(c *gin.Context) {
	id, err := shared.ConvertStringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "ID inválido", nil))
		return
	}

	useCase := usecase.NewEliminarCursoUseCase(di.GetContainer().GetCursoRepository())
	response := useCase.Execute(id)
	c.JSON(response.StatusCode, response)
}
