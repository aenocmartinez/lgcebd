package controller

import (
	"ebd/src/infraestructure/di"
	"ebd/src/shared"
	usecase "ebd/src/usecase/clase"
	"ebd/src/view/dto"
	formrequest "ebd/src/view/formrequest/clase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegistrarAsistencia(c *gin.Context) {
	var req formrequest.RegistrarAsitenciaFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewAPIResponse(http.StatusBadRequest, "Datos de entrada inválidos", err.Error()))
		return
	}

	datos := dto.GuardarClaseDTO{
		Fecha:               req.Fecha,
		Ofrenda:             req.Ofrenda,
		GrupoID:             req.GrupoID,
		ContenidoTematicoID: req.ContenidoTematicoID,
	}

	registrarAsitencia := usecase.NewRegistrarAsistenciaUseCase(
		di.GetContainer().GetClaseRepository(),
		di.GetContainer().GetGrupoRepository(),
		di.GetContainer().GetContenidoTematicoRepository())

	response := registrarAsitencia.Execute(datos)

	c.JSON(response.StatusCode, response)
}
