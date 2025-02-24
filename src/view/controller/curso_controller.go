package controller

import (
	"ebd/src/usecase/curso"

	"github.com/gin-gonic/gin"
)

func ListarCursos(c *gin.Context) {

	listarCursosUseCase := curso.ListarCursosUseCase{}

	response := listarCursosUseCase.Execute()

	c.JSON(response.StatusCode, response)
}
